# -*- coding: UTF-8 -*-
'''
@Project ：FuChuang 
@File    ：web.py
@Author  ：Max
@Date    ：2024-03-25 21:58 
@note    :
'''


from flask import Flask, request, jsonify

import json_parse

app=Flask(__name__)

@app.route('/json/parse',methods=['POST'])
def ParseJson():

    file_path=request.json.get('file_path')
    stay_time=request.json.get('staytime')
    repeat_click=request.json.get('repeat_click')
    page_load=request.json.get('page_load')
    feedback_interval=request.json.get('feedback_interval')
    no_reaction=request.json.get('no_reaction')
    error_count=request.json.get('error_count')
    console_errors=request.json.get('console_errors')
    occur_many=request.json.get('occur_many')

    weights=[stay_time, repeat_click, page_load, feedback_interval, no_reaction,error_count,console_errors,occur_many]


    res = json_parse.jsonParse(file_path, weights)


    #todo 补全一下这里的响应
    resp={
        "score":res.m_score
    }

    return jsonify(resp)


