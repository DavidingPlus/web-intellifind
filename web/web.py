# -*- coding: UTF-8 -*-
'''
@Project ：FuChuang 
@File    ：web.py
@Author  ：Max
@Date    ：2024-03-25 21:58 
@note    :
'''


from flask import Flask, request, jsonify
from web.json_parse import jsonParse

app = Flask(__name__)


@app.route('/json/parse', methods=['POST'])
def ParseJson():
    data = request.get_json()
    save_path=data['save_path']

    print(save_path)
    stay_time = data['stay_time']
    repeat_click = data['repeat_click']
    page_load =data['page_load']
    feedback = data['feedback']
    no_reaction = data['no_reaction']
    error_count = data['error_count']
    console_errors = data['console_errors']
    is_blank = data['is_blank']
    occur_many =data['occur_many']

    weights = [stay_time, repeat_click, page_load, feedback,
               no_reaction, error_count, console_errors, is_blank,occur_many]

    print(weights)
    res = jsonParse(save_path, weights)

    print(res)
    resp = {
        # 总分
        "total_score": res.m_totalScore,
        # 各个部分的得分
        "stay_time": res.m_stayTimeScore,
        "repeat_click": res.m_repeatClickScore,
        "page_load": res.m_pageLoadScore,
        "feedback_interval": res.m_feedbackIntervalScore,
        "no_reaction": res.m_noReactionScore,
        "error_count": res.m_errorCountScore,
        "console_errors": res.m_consoleErrorsScore,
        "is_blank": res.m_isBlankScore,
        "occur_many": res.m_occurManyScore,
        # 简要描述
        "brief_desc": res.m_briefDesc,
        # 详细描述
        "detail_desc": res.m_detailDesc
    }

    return jsonify(resp)
