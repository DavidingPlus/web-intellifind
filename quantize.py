import json
import random
import torch

from auto_gptq import AutoGPTQForCausalLM, BaseQuantizeConfig
from datasets import Dataset
from transformers import AutoTokenizer


def load_data(data_path, tokenizer, n_samples):
    with open(data_path, "r", encoding="utf-8") as f:
        raw_data = json.load(f)

    raw_data = random.sample(raw_data, k=min(n_samples, len(raw_data)))

    def dummy_gen():
        return raw_data

    def tokenize(examples):
        instructions = examples["instruction"]
        inputs = examples["input"]
        outputs = examples["output"]

        prompts = []
        texts = []
        input_ids = []
        attention_mask = []
        for istr, inp, opt in zip(instructions, inputs, outputs):
            if inp:
                prompt = f"Instruction:\n{istr}\nInput:\n{inp}\nOutput:\n"
                text = prompt + opt
            else:
                prompt = f"Instruction:\n{istr}\nOutput:\n"
                text = prompt + opt
            if len(tokenizer(prompt)["input_ids"]) >= tokenizer.model_max_length:
                continue

            tokenized_data = tokenizer(text)

            input_ids.append(
                tokenized_data["input_ids"][: tokenizer.model_max_length])
            attention_mask.append(
                tokenized_data["attention_mask"][: tokenizer.model_max_length])
            prompts.append(prompt)
            texts.append(text)

        return {
            "input_ids": input_ids,
            "attention_mask": attention_mask,
            "prompt": prompts,
        }

    dataset = Dataset.from_generator(dummy_gen)

    dataset = dataset.map(
        tokenize,
        batched=True,
        batch_size=len(dataset),
        num_proc=1,
        keep_in_memory=True,
        load_from_cache_file=False,
        remove_columns=["instruction", "input"],
    )

    dataset = dataset.to_list()

    for sample in dataset:
        sample["input_ids"] = torch.LongTensor(sample["input_ids"])
        sample["attention_mask"] = torch.LongTensor(sample["attention_mask"])

    return dataset


def main():
    pretrained_model_dir = "../Llama2-Chinese-13b-Chat"
    quantized_model_dir = "../Llama2-Chinese-13b-Chat-quantized"
    dataset_dir = "../../alpaca_data_cleaned.json"
    n_samples = 1024

    quantize_config = BaseQuantizeConfig(
        bits=4,  # 将模型量化为 4-bit 数值类型
        group_size=128,  # 一般推荐将此参数的值设置为 128
        desc_act=False,  # 设为 False 可以显著提升推理速度，但是 ppl 可能会轻微地变差
    )

    # 加载未量化的模型，默认情况下，模型总是会被加载到 CPU 内存中
    model = AutoGPTQForCausalLM.from_pretrained(
        pretrained_model_dir, quantize_config)
    tokenizer = AutoTokenizer.from_pretrained(pretrained_model_dir)
    examples = load_data(dataset_dir, tokenizer, n_samples)
    examples_for_quant = [
        {"input_ids": example["input_ids"], "attention_mask": example["attention_mask"]} for example in examples
    ]

    model.quantize(examples_for_quant)
    model.save_quantized(quantized_model_dir)
    model.save_quantized(quantized_model_dir, use_safetensors=True)


if __name__ == '__main__':
    main()
