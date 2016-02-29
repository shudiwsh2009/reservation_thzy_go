var width=$(window).width();
var height=$(window).height();
var teacher;
var reservations;
var knowingMethods;

function viewReservations() {
  getKnowingMethods();
  $.getJSON('/teacher/reservation/view', function(json, textStatus) {
    if (json.state === 'SUCCESS') {
      console.log(json);
      reservations = json.reservations;
      teacher = json.teacher_info;
      refreshDataTable(reservations);
      optimize();
    } else {
      alert(json.message);
    }
  });
}

function queryStudent() {
  $.post('/admin/student/query', {
    student_username: $('#query_student').val()
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      showStudent(data.student_info);
    } else {
      alert(data.message);
    }
  });
}

function getKnowingMethods() {
  $.getJSON('/user/knowing_methods', function(json, textStatus) {
    if (json.state === 'SUCCESS') {
      knowingMethods = json.knowing_methods;
    }
  });
}

function refreshDataTable(reservations) {
  $('#page_maintable')[0].innerHTML = '\
    <div class="table_col" id="col_select">\
      <div class="table_head table_cell" id="head_select">\
        <button id="btn_select_all" name="all" onclick="selectAll();" style="padding: 0px;">全选</button>\
      </div>\
    </div>\
    <div class="table_col" id="col_time">\
      <div class="table_head table_cell">时间</div>\
    </div>\
    <div class="table_col" id="col_teacher_fullname">\
      <div class="table_head table_cell">咨询师</div>\
    </div>\
    <div class="table_col" id="col_teacher_mobile">\
      <div class="table_head table_cell">咨询师手机</div>\
    </div>\
    <div class="table_col" id="col_status">\
      <div class="table_head table_cell">状态</div>\
    </div>\
    <div class="table_col" id="col_student">\
      <div class="table_head table_cell">学生</div>\
    </div>\
    <div class="clearfix"></div>\
  ';

  for (var i = 0; i < reservations.length; ++i) {
    $('#col_select').append('<div class="table_cell" id="cell_select_' + i + '">'
      + '<input class="checkbox" type="checkbox" id="cell_checkbox_' + i + '"></div>');
    $('#col_time').append('<div class="table_cell" id="cell_time_' + i + '" name="' + i + '">' + reservations[i].start_time.split(' ')[0].substr(2) + '<br>' 
      + reservations[i].start_time.split(' ')[1] + '-' + reservations[i].end_time.split(' ')[1] + '</div>');
    $('#col_teacher_fullname').append('<div class="table_cell" id="cell_teacher_fullname_'
      + i + '">' + reservations[i].teacher_fullname + '</div>');
    $('#col_teacher_mobile').append('<div class="table_cell" id="cell_teacher_mobile_'
      + i + '">' + reservations[i].teacher_mobile + '</div>');
    if (reservations[i].status === 'AVAILABLE') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i + '">未预约</div>');
      $('#col_student').append('<div class="table_cell" id="cell_student_' + i + '">' 
        + '<button type="button" id="cell_student_view_' + i + '" disabled="true" style="padding: 2px 2px">查看'
        + '</button></div>');
    } else if (reservations[i].status === 'RESERVATED') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i + '">已预约</div>');
      $('#col_student').append('<div class="table_cell" id="cell_student_' + i + '">' 
        + '<button type="button" id="cell_student_view_' + i + '" onclick="getStudent(' + i + ');" style="padding: 2px 2px">查看'
        + '</button></div>');
    } else if (reservations[i].status === 'FEEDBACK') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i + '">'
        + '<button type="button" id="cell_status_feedback_' + i + '" onclick="getFeedback(' + i + ');" style="padding: 2px 2px">'
        + '反馈</button></div>');
      $('#col_student').append('<div class="table_cell" id="cell_student_' + i + '">' 
        + '<button type="button" id="cell_student_view_' + i + '" onclick="getStudent(' + i + ');" style="padding: 2px 2px">查看'
        + '</button></div>');
    }
  }
  $('#col_select').append('<div class="table_cell" id="cell_select_add"><input type="checkbox"></div>');
  $('#col_time').append('<div class="table_cell" id="cell_time_add">点击新增</div>');
  $('#col_teacher_fullname').append('<div class="table_cell" id="cell_teacher_fullname_add"></div>');
  $('#col_teacher_mobile').append('<div class="table_cell" id="cell_teacher_mobile_add"></div>');
  $('#col_status').append('<div class="table_cell" id="cell_status_add"></div>');
  $('#col_student').append('<div class="table_cell" id="cell_student_add"></div>');
  $(function() {
    for (var i = 0; i < reservations.length; i++) {
      $('#cell_time_' + i).click(function(e) {
        editReservation($(e.target).attr('name'));
      });
    }
    $('#cell_time_add').click(addReservation);
  });
}

function selectAll() {
  if ($('#btn_select_all').prop('name') && $('#btn_select_all').prop('name') === 'all') {
    $('.checkbox').prop('checked', true);
    $('#btn_select_all').prop('name', 'none');
    $('#btn_select_all').text('不选');
  } else {
    $('.checkbox').prop('checked', false);
    $('#btn_select_all').prop('name', 'all');
    $('#btn_select_all').text('全选');
  }
}

function optimize(t) {
  $("#col_select").width(width * 0.08);
  $('#col_time').width(width * 0.23);
  $('#col_teacher_fullname').width(width * 0.15);
  $('#col_teacher_mobile').width(width * 0.24);
  $('#col_status').width(width * 0.12);
  $('#col_student').width(width * 0.12);
  $('#col_select').css('margin-left', width * 0.01 + 'px');
  for (var i = 0; i < reservations.length; ++i) {
    var maxHeight = Math.max(
        $('#cell_select_' + i).height(),
        $('#cell_time_' + i).height(),
        $('#cell_teacher_fullname_' + i).height(),
        $('#cell_teacher_mobile_' + i).height(),
        $('#cell_status_' + i).height(),
        $('#cell_student_' + i).height()
      );
    $('#cell_select_' + i).height(maxHeight);
    $('#cell_time_' + i).height(maxHeight);
    $('#cell_teacher_fullname_' + i).height(maxHeight);
    $('#cell_teacher_mobile_' + i).height(maxHeight);
    $('#cell_status_' + i).height(maxHeight);
    $('#cell_student_' + i).height(maxHeight);

    if (i % 2 == 1) {
      $('#cell_select_' + i).css('background-color', 'white');
      $('#cell_time_' + i).css('background-color', 'white');
      $('#cell_teacher_fullname_' + i).css('background-color', 'white');
      $('#cell_teacher_mobile_' + i).css('background-color', 'white');
      $('#cell_status_' + i).css('background-color', 'white');
      $('#cell_student_' + i).css('background-color', 'white');
    } else {
      $('#cell_select_' + i).css('background-color', '#f3f3ff');
      $('#cell_time_' + i).css('background-color', '#f3f3ff');
      $('#cell_teacher_fullname_' + i).css('background-color', '#f3f3ff');
      $('#cell_teacher_mobile_' + i).css('background-color', '#f3f3ff');
      $('#cell_status_' + i).css('background-color', '#f3f3ff');
      $('#cell_student_' + i).css('background-color', '#f3f3ff');
    }
  }
  var s = 28;
  if (t === "add") {
    s = 68;
  }
  $("#cell_select_add").height(s);
  $("#cell_time_add").height(s);
  $("#cell_teacher_fullname_add").height(s);
  $("#cell_teacher_mobile_add").height(s);
  $("#cell_status_add").height(s);
  $("#cell_student_add").height(s);

  $('.table_head').height($('#head_select').height());
  $(t).css('left', (width - $(t).width()) / 2 - 11 + 'px');
  $(t).css('top', (height - $(t).height()) / 2 - 11 + 'px');
}

function addReservation() {
  $('#cell_time_add').off();
  $('#cell_time_add').first().html('<input type="text" id="input_date_add" style="width: 60px;"/><br>'
    + '<input style="width:15px;" id="start_hour_add"/>时<input style="width:15px" id="start_minute_add"/>分<br>'
    + '<input style="width:15px;" id="end_hour_add"/>时<input style="width:15px" id="end_minute_add"/>分');
  $('#cell_teacher_fullname_add').first().html('<input id="teacher_fullname_add" style="width:80px" value="' + teacher.teacher_fullname + '"/>');
  $('#cell_teacher_mobile_add').first().html('<input id="teacher_mobile_add" style="width:120px" value="' + teacher.teacher_mobile + '"/>');
  $('#cell_status_add').first().html('<button type="button" onclick="addReservationConfirm();">确认</button>');
  $('#cell_student_add').first().html('<button type="button" onclick="window.location.reload();">取消</button>');
  $('#input_date_add').datepicker({
    showOtherMonths: true,
    selectOtherMonths: true,
    showButtonPanel: true,
    dateFormat: 'yy-mm-dd',
    showWeek: true,
    firstDay: 1
  });
  optimize("add");
}

function addReservationConfirm() {
  var startHour = $('#start_hour_add').val();
  var startMinute = $('#start_minute_add').val();
  var endHour = $('#end_hour_add').val();
  var endMinute = $('#end_minute_add').val();
  var startTime = $('#input_date_add').val() + ' ' + (startHour.length < 2 ? '0' : '') + startHour + ':';
  if (startMinute.length == 0) {
    startTime += '00';
  } else if (startMinute.length == 1) {
    startTime += '0' + startMinute;
  } else {
    startTime += startMinute;
  }
  var endTime = $('#input_date_add').val() + ' ' + (endHour.length < 2 ? '0' : '') + endHour + ':';
  if (endMinute.length == 0) {
    endTime += '00';
  } else if (endMinute.length == 1) {
    endTime += '0' + endMinute;
  } else {
    endTime += endMinute;
  }
  var payload = {
    start_time: startTime,
    end_time: endTime,
    teacher_fullname: $('#teacher_fullname_add').val(),
    teacher_mobile: $('#teacher_mobile_add').val(),
  };
  $.post('/teacher/reservation/add', payload, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else if (data.state === 'CHECK') {
      addReservationCheck(payload);
    } else {
      alert(data.message);
    }
  });
}

function addReservationCheck(payload) {
  $('body').append('\
    <div id="pop_add_reservation_check" class="pop_window" style="width: 50%">\
      咨询师信息有变更，是否更新？\
      <br>\
      <button type="button" name="confirm">确认</button>\
      <button type="button" onclick="$(\'#pop_add_reservation_check\').remove();">取消</button>\
    </div>\
  ');
  $(function() {
    $('#pop_add_reservation_check [name=confirm]').click(function() {
      $('#pop_add_reservation_check').remove();
      addReservationCheckConfirm(payload);
    });
  });
  optimize('#pop_add_reservation_check');
}

function addReservationCheckConfirm(payload) {
  payload['force'] = 'FORCE';
  $.post('/teacher/reservation/add', payload, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else if (data.state === 'CHECK') {
      addReservationCheck(payload);
    } else {
      alert(data.message);
    }
  });
}

function editReservation(index) {
  $("#cell_time_" + index).height(68);
  $('#cell_time_' + index).off();
  $('#cell_time_' + index).first().html('<input type="text" id="input_date_' + index + '" style="width: 60px;"/><br>'
    + '<input style="width:15px;" id="start_hour_' + index + '"/>时<input style="width:15px" id="start_minute_' + index + '"/>分<br>'
    + '<input style="width:15px;" id="end_hour_' + index + '"/>时<input style="width:15px" id="end_minute_' + index + '"/>分');
  $('#cell_teacher_fullname_' + index).first().html('<input id="teacher_fullname_' + index + '" style="width:80px" value="' + teacher.teacher_fullname + '"/>');
  $('#cell_teacher_mobile_' + index).first().html('<input id="teacher_mobile_' + index + '" style="width:120px" value="' + teacher.teacher_mobile + '"/>');
  $('#cell_status_' + index).first().html('<button type="button" onclick="editReservationConfirm(' + index + ');">确认</button>');
  $('#cell_student_' + index).first().html('<button type="button" onclick="window.location.reload();">取消</button>');
  $('#input_date_' + index).datepicker({
    showOtherMonths: true,
    selectOtherMonths: true,
    showButtonPanel: true,
    dateFormat: 'yy-mm-dd',
    showWeek: true,
    firstDay: 1
  });
  optimize();
}

function editReservationConfirm(index) {
  var startHour = $('#start_hour_' + index).val();
  var startMinute = $('#start_minute_' + index).val();
  var endHour = $('#end_hour_' + index).val();
  var endMinute = $('#end_minute_' + index).val();
  var startTime = $('#input_date_' + index).val() + ' ' + (startHour.length < 2 ? '0' : '') + startHour + ':';
  if (startMinute.length == 0) {
    startTime += '00';
  } else if (startMinute.length == 1) {
    startTime += '0' + startMinute;
  } else {
    startTime += startMinute;
  }
  var endTime = $('#input_date_' + index).val() + ' ' + (endHour.length < 2 ? '0' : '') + endHour + ':';
  if (endMinute.length == 0) {
    endTime += '00';
  } else if (endMinute.length == 1) {
    endTime += '0' + endMinute;
  } else {
    endTime += endMinute;
  }
  var payload = {
    reservation_id: reservations[index].reservation_id,
    start_time: startTime,
    end_time: endTime,
    teacher_fullname: $('#teacher_fullname_' + index).val(),
    teacher_mobile: $('#teacher_mobile_' + index).val(),
  };
  $.post('/teacher/reservation/edit', payload, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else if (data.state === 'CHECK') {
      editReservationCheck(payload);
    } else {
      alert(data.message);
    }
  });
}

function editReservationCheck(payload) {
  $('body').append('\
    <div id="pop_edit_reservation_check" class="pop_window" style="width: 50%">\
      咨询师信息有变更，是否更新？\
      <br>\
      <button type="button" name="confirm">确认</button>\
      <button type="button" onclick="$(\'#pop_edit_reservation_check\').remove();">取消</button>\
    </div>\
  ');
  $(function() {
    $('#pop_edit_reservation_check [name=confirm]').click(function() {
      $('#pop_edit_reservation_check').remove();
      editReservationCheckConfirm(payload);
    });
  });
  optimize('#pop_edit_reservation_check');
}

function editReservationCheckConfirm(payload) {
  payload['force'] = 'FORCE';
  $.post('/teacher/reservation/edit', payload, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else if (data.state === 'CHECK') {
      addReservationCheck(payload);
    } else {
      alert(data.message);
    }
  });
}

function removeReservations() {
  $('body').append('\
    <div id="pop_remove_reservations" class="pop_window" style="width: 50%">\
      确认删除选中的咨询记录？\
      <br>\
      <button type="button" onclick="$(\'#pop_remove_reservations\').remove();removeReservationsConfirm();">确认</button>\
      <button type="button" onclick="$(\'#pop_remove_reservations\').remove();">取消</button>\
    </div>\
  ');
  optimize('#pop_remove_reservations');
}

function removeReservationsConfirm() {
  var reservationIds = [];
  for (var i = 0; i < reservations.length; ++i) {
    if ($('#cell_checkbox_' + i)[0].checked) {
      reservationIds.push(reservations[i].reservation_id);
    }
  }
  var payload = {
    reservation_ids: reservationIds,
  };
  $.ajax({
    url: '/teacher/reservation/remove',
    type: "POST",
    dataType: 'json',
    data: payload,
    traditional: true,
  })
  .done(function(data) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else {
      alert(data.message);
    }
  });
}

function cancelReservations() {
  $('body').append('\
    <div id="pop_cancel_reservations" class="pop_window" style="width: 50%">\
      确认取消选中的预约？\
      <br>\
      <button type="button" onclick="$(\'#pop_cancel_reservations\').remove();cancelReservationsConfirm();">确认</button>\
      <button type="button" onclick="$(\'#pop_cancel_reservations\').remove();">取消</button>\
    </div>\
  ');
  optimize('#pop_cancel_reservations');
}

function cancelReservationsConfirm() {
  var reservationIds = [];
  for (var i = 0; i < reservations.length; ++i) {
    if ($('#cell_checkbox_' + i)[0].checked) {
      reservationIds.push(reservations[i].reservation_id);
    }
  }
  var payload = {
    reservation_ids: reservationIds,
  };
  $.ajax({
    url: '/teacher/reservation/cancel',
    type: "POST",
    dataType: 'json',
    data: payload,
    traditional: true,
  })
  .done(function(data) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else {
      alert(data.message);
    }
  });
}

function getStudent(index) {
  $.post('/teacher/student/get', {
    student_id: reservations[index].student_id
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      showStudent(data.student_info, data.reservations);
    } else {
      alert(data.message);
    }
  });
}

function showStudent(student, reservations) {
  $('body').append('\
    <div id="pop_show_student_' + student.student_id + '" class="pop_window" style="text-align: left; width: 90%; height: 60%; overflow: auto">\
      学号：' + student.student_username + '<br>\
      姓名：' + student.student_fullname + '<br>\
      性别：' + student.student_gender + '<br>\
      出生日期：' + student.student_birthday + '<br>\
      系别：' + student.student_school + '<br>\
      年级：' + student.student_grade + '<br>\
      现住址：' + student.student_current_address + '<br>\
      家庭住址：' + student.student_family_address + '<br>\
      联系电话：' + student.student_mobile + '<br>\
      Email：' + student.student_email + '<br>\
      咨询经历：' + (student.student_experience_time ? '时间：' + student.student_experience_time + ' 地点：' + student.student_experience_location + ' 咨询师：' + student.student_experience_teacher : '无') + '<br>\
      父亲年龄：' + student.student_father_age + ' 职业：' + student.student_father_job + ' 学历：' + student.student_father_edu + '<br>\
      母亲年龄：' + student.student_mother_age + ' 职业：' + student.student_mother_job + ' 学历：' + student.student_mother_edu + '<br>\
      父母婚姻状况：' + student.student_parent_marriage + '<br>\
      近三个月里发生的有重大意义的事：' + student.student_significant + '<br>\
      需要接受帮助的主要问题：' + student.student_problem + '<br>\
      档案分类：' + student.student_archive_category + ' 档案编号：' + student.student_archive_number + '<br>\
      是否危机个案：<span id="crisis_level_'+ student.student_id + '"></span><br>\
      <div id="key_case_' + student.student_id + '" style="display: none">\
        <b>重点个案：</b>\
        <input id="key_case_' + student.student_id + '_0" type="checkbox">通报院系</input>\
        <input id="key_case_' + student.student_id + '_1" type="checkbox">联席会议</input>\
        <input id="key_case_' + student.student_id + '_2" type="checkbox">服药</input>\
        <input id="key_case_' + student.student_id + '_3" type="checkbox">自杀未遂</input>\
        <input id="key_case_' + student.student_id + '_4" type="checkbox">家长陪读</input>\
        <br>\
        <b>医疗诊断：</b>\
        <input id="medical_diagnosis_' + student.student_id + '_0" type="checkbox">精神分裂诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_1" type="checkbox">双相诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_2" type="checkbox">抑郁症诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_3" type="checkbox">强迫症诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_4" type="checkbox">进食障碍诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_5" type="checkbox">失眠诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_6" type="checkbox">其他精神症状诊断</input>\
        <input id="medical_diagnosis_' + student.student_id + '_7" type="checkbox">躯体疾病诊断</input>\
      </div>\
      <br>\
      已绑定的咨询师：<span id="binded_teacher_username">' + student.student_binded_teacher_username + '</span>&nbsp;\
        <span id="binded_teacher_fullname">' + student.student_binded_teacher_fullname + '</span><br>\
      <br>\
      <button type="button" onclick="$(\'#pop_show_student_' + student.student_id + '\').remove();">关闭</button>\
      <div id="student_reservations_' + student.student_id + '" style="width: 100%">\
      </div>\
    </div>\
  ');
  for (var i = 0; i < reservations.length; i++) {
    $('#student_reservations_' + student.student_id).append('\
      <div class="has_children" style="background: ' + (reservations[i].status === 'FEEDBACK' ? '#555' : '#F00') + '">\
        <span>' + reservations[i].start_time + '  ' + reservations[i].teacher_fullname + '</span>\
        <p class="children" style="width:100%;">学生反馈：' + reservations[i].student_feedback.scores + '</p>\
        <p class="children" style="width:100%;">评估分类：' + reservations[i].teacher_feedback.category + '</p>\
        <p class="children" style="width:100%;">出席人员：' + reservations[i].teacher_feedback.participants + '</p>\
        <p class="children" style="width:100%;">问题评估：' + reservations[i].teacher_feedback.problem + '</p>\
        <p class="children" style="width:100%;">咨询记录：' + reservations[i].teacher_feedback.record + '</p>\
      </div>\
    ');
  }
  $(function() {
    $('.has_children').click(function() {
      $(this).addClass('highlight').children('p').show().end()
          .siblings().removeClass('highlight').children('p').hide();
    });
    var i = 1;
    for (i = 0; i < 5; i++) {
      $('#key_case_' + student.student_id + '_' + i).first().attr('checked', student.student_key_case[i] > 0);
    }
    for (i = 0; i < 8; i++) {
      $('#medical_diagnosis_' + student.student_id + '_' + i).first().attr('checked', student.student_medical_diagnosis[i] > 0);
    }
    $('#crisis_level_' + student.student_id).text(student.student_crisis_level === 0 ? '否' : '是');
    if (student.student_crisis_level !== 0) {
      $('#key_case_' + student.student_id).show();
    }
  });
  optimize('#pop_show_student_' + student.student_id);
}