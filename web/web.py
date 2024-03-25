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
    file_path = request.json.get('file_path')
    stay_time = request.json.get('staytime')
    repeat_click = request.json.get('repeat_click')
    page_load = request.json.get('page_load')
    feedback_interval = request.json.get('feedback_interval')
    no_reaction = request.json.get('no_reaction')
    error_count = request.json.get('error_count')
    console_errors = request.json.get('console_errors')
    occur_many = request.json.get('occur_many')

    weights = [stay_time, repeat_click, page_load, feedback_interval,
               no_reaction, error_count, console_errors, occur_many]

    res = jsonParse(file_path, weights)

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
