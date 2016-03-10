var width = $(window).width();
var height = $(window).height();
var reservations;
var knowingMethods;

function viewReservations() {
  getKnowingMethods();
  if ($('#query_date').val() !== '') {
    queryReservations();
    return;
  }
  $.getJSON('/admin/reservation/view', function(json, textStatus) {
    if (json.state === 'SUCCESS') {
      console.log(json);
      reservations = json.reservations;
      refreshDataTable(reservations);
      optimize();
    } else {
      alert(json.message);
    }
  });
}

function queryReservations() {
  $.getJSON('/admin/reservation/view/monthly', {
    from_date: $('#query_date').val()
  }, function(json, textStatus) {
    if (json.state === 'SUCCESS') {
      console.log(json);
      reservations = json.reservations;
      refreshDataTable(reservations);
      optimize();
    } else {
      alert(json.message);
    }
  });
}

// function exportTodayReservations() {
//   $.getJSON('/admin/reservation/export/today', function(json, textStatus) {
//     if (json.state === 'SUCCESS') {
//       window.open(json.url);
//     } else {
//       alert(json.message);
//     }
//   });
// }

function exportReservations() {
  var reservationIds = [];
  for (var i = 0; i < reservations.length; ++i) {
    if ($("#cell_checkbox_" + i)[0].checked) {
      reservationIds.push(reservations[i].reservation_id);
    }
  }
  var payload = {
    reservation_ids: reservationIds,
  };
  $.ajax({
    type: "POST",
    url: "/admin/reservation/export",
    data: payload,
    traditional: true,
    dataType: "json",
    success: function(data) {
      if (data.state === "SUCCESS") {
        window.open(data.url);
      } else {
        alert(data.message);
      }
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
  $('#page_maintable').first().html('\
    <div class="table_col" id="col_select">\
      <div class="table_head table_cell" id="head_select">\
        <button id="btn_select_all" name="all" onclick="selectAll();">全选</button>\
      </div>\
    </div>\
    <div class="table_col" id="col_time">\
      <div class="table_head table_cell">时间</div>\
    </div>\
    <div class="table_col" id="col_teacher_fullname">\
      <div class="table_head table_cell">咨询师</div>\
    </div>\
    <div class="table_col" id="col_teacher_username">\
      <div class="table_head table_cell">咨询师编号</div>\
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
  ');

  for (var i = 0; i < reservations.length; ++i) {
    $('#col_select').append('<div class="table_cell" id="cell_select_' + i + '">'
      + '<input class="checkbox" type="checkbox" id="cell_checkbox_' + i + '"></div>');
    $('#col_time').append('<div class="table_cell" id="cell_time_' + i + '" name="' + i + '">' + reservations[i].start_time + 
      ' 至 ' + reservations[i].end_time + '</div>');
    $('#col_teacher_fullname').append('<div class="table_cell" id="cell_teacher_fullname_'
      + i + '">' + reservations[i].teacher_fullname + '</div>');
    $('#col_teacher_username').append('<div class="table_cell" id="cell_teacher_username_'
      + i + '">' + reservations[i].teacher_username + '</div>');
    $('#col_teacher_mobile').append('<div class="table_cell" id="cell_teacher_mobile_'
      + i + '">' + reservations[i].teacher_mobile + '</div>');
    if (reservations[i].status === 'AVAILABLE') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i + '">未预约</div>');
      $('#col_student').append('<div class="table_cell" id="cell_student_' + i + '">' 
        + '<button type="button" id="cell_student_view_' + i + '" onclick="setStudent(' + i + ')">指定'
        + '</button></div>');
    } else if (reservations[i].status === 'RESERVATED') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i + '">已预约</div>');
      $('#col_student').append('<div class="table_cell" id="cell_student_' + i + '">' 
        + '<button type="button" id="cell_student_view_' + i + '" onclick="getStudent(' + i + ');">查看'
        + '</button></div>');
    } else if (reservations[i].status === 'FEEDBACK') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i + '">已预约</div>');
      $('#col_student').append('<div class="table_cell" id="cell_student_' + i + '">' 
        + '<button type="button" id="cell_student_view_' + i + '" onclick="getStudent(' + i + ');">查看'
        + '</button></div>');
    }
  }
  $('#col_select').append('<div class="table_cell" id="cell_select_add"><input type="checkbox"></div>');
  $('#col_time').append('<div class="table_cell" id="cell_time_add">点击新增</div>');
  $('#col_teacher_fullname').append('<div class="table_cell" id="cell_teacher_fullname_add"></div>');
  $('#col_teacher_username').append('<div class="table_cell" id="cell_teacher_username_add"></div>');
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
  $('#col_select').width(40);
  $('#col_time').width(405);
  $('#col_teacher_fullname').width(120);
  $('#col_teacher_username').width(160);
  $('#col_teacher_mobile').width(160);
  $('#col_status').width(85);
  $('#col_student').width(85);
  for (var i = 0; i < reservations.length; ++i) {
    var maxHeight = Math.max(
      $('#cell_select_' + i).height(),
      $('#cell_time_' + i).height(),
      $('#cell_teacher_fullname_' + i).height(),
      $('#cell_teacher_username_' + i).height(),
      $('#cell_teacher_mobile_' + i).height(),
      $('#cell_status_' + i).height(),
      $('#cell_student_' + i).height()
    );
    $('#cell_select_' + i).height(maxHeight);
    $('#cell_time_' + i).height(maxHeight);
    $('#cell_teacher_fullname_' + i).height(maxHeight);
    $('#cell_teacher_username_' + i).height(maxHeight);
    $('#cell_teacher_mobile_' + i).height(maxHeight);
    $('#cell_status_' + i).height(maxHeight);
    $('#cell_student_' + i).height(maxHeight);

    if (i % 2 == 1) {
      $('#cell_select_' + i).css('background-color', 'white');
      $('#cell_time_' + i).css('background-color', 'white');
      $('#cell_teacher_fullname_' + i).css('background-color', 'white');
      $('#cell_teacher_username_' + i).css('background-color', 'white');
      $('#cell_teacher_mobile_' + i).css('background-color', 'white');
      $('#cell_status_' + i).css('background-color', 'white');
      $('#cell_student_' + i).css('background-color', 'white');
    } else {
      $('#cell_select_' + i).css('background-color', '#f3f3ff');
      $('#cell_time_' + i).css('background-color', '#f3f3ff');
      $('#cell_teacher_fullname_' + i).css('background-color', '#f3f3ff');
      $('#cell_teacher_username_' + i).css('background-color', '#f3f3ff');
      $('#cell_teacher_mobile_' + i).css('background-color', '#f3f3ff');
      $('#cell_status_' + i).css('background-color', '#f3f3ff');
      $('#cell_student_' + i).css('background-color', '#f3f3ff');
    }

    if (reservations[i].student_crisis_level && reservations[i].student_crisis_level !== '0') {
      $('#cell_student_' + i).css('background-color', 'rgba(255, 0, 0, ' + parseInt(reservations[i].student_crisis_level) / 1 +')');
    }
  }
  $('#cell_select_add').height(28);
  $('#cell_time_add').height(28);
  $('#cell_teacher_fullname_add').height(28);
  $('#cell_teacher_username_add').height(28);
  $('#cell_teacher_mobile_add').height(28);
  $('#cell_status_add').height(28);
  $('#cell_student_add').height(28);

  $('.table_head').height($('#head_select').height());
  $(t).css('left', (width - $(t).width()) / 2 - 11 + 'px');
  $(t).css('top', (height - $(t).height()) / 2 - 11 + 'px');
  $('#page_maintable').css('margin-left', 0.5 * ($(window).width()
    - (40 + 405 + 120 + 85 + 85 + 320)) + 'px');
}

function addReservation() {
  $('#cell_time_add').off();
  $('#cell_time_add').first().html('<input type="text" id="input_date_add" style="width: 80px"/>日，'
    + '<input style="width:20px" id="start_hour_add"/>时<input style="width:20px" id="start_minute_add"/>分'
    + '至<input style="width:20px" id="end_hour_add"/>时<input style="width:20px" id="end_minute_add"/>分');
  $('#cell_teacher_fullname_add').first().html('<input id="teacher_fullname_add" style="width:60px"/>'
    + '<button type="button" onclick="searchTeacher();">搜索</button>');
  $('#cell_teacher_username_add').first().html('<input id="teacher_username_add" style="width:120px"/>');
  $('#cell_teacher_mobile_add').first().html('<input id="teacher_mobile_add" style="width:120px"/>');
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
  optimize();
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
    teacher_username: $('#teacher_username_add').val(),
    teacher_fullname: $('#teacher_fullname_add').val(),
    teacher_mobile: $('#teacher_mobile_add').val(),
  };
  $.post('/admin/reservation/add', payload, function(data, textStatus, xhr) {
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
  $.post('/admin/reservation/add', payload, function(data, textStatus, xhr) {
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
  $('#cell_time_' + index).off();
  $('#cell_time_' + index).first().html('<input type="text" id="input_date_' + index + '" style="width: 80px"/>日，'
    + '<input style="width:20px" id="start_hour_' + index + '"/>时<input style="width:20px" id="start_minute_' + index + '"/>分'
    + '至<input style="width:20px" id="end_hour_' + index + '"/>时<input style="width:20px" id="end_minute_' + index + '"/>分');
  $('#cell_teacher_fullname_' + index).first().html('<input id="teacher_fullname_' + index + '" style="width:60px" '
    + 'value="' + reservations[index].teacher_fullname + '""></input>'
    + '<button type="button" onclick="searchTeacher(' + index + ');">搜索</button>');
  $('#cell_teacher_username_' + index).first().html('<input id="teacher_username_' + index + '" style="width:120px" '
    + 'value="' + reservations[index].teacher_username + '"/>');
  $('#cell_teacher_mobile_' + index).first().html('<input id="teacher_mobile_' + index + '" style="width:120px" '
    + 'value="' + reservations[index].teacher_mobile + '"/>');
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
    teacher_username: $('#teacher_username_' + index).val(),
    teacher_fullname: $('#teacher_fullname_' + index).val(),
    teacher_mobile: $('#teacher_mobile_' + index).val(),
  };
  $.post('/admin/reservation/edit', payload, function(data, textStatus, xhr) {
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
  $.post('/admin/reservation/edit', payload, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      viewReservations();
    } else if (data.state === 'CHECK') {
      addReservationCheck(payload);
    } else {
      alert(data.message);
    }
  });
}

function searchTeacher(index) {
  $.post('/admin/teacher/search', {
    teacher_username: $('#teacher_username_' + (index === undefined ? 'add' : index)).val(),
    teacher_fullname: $('#teacher_fullname_' + (index === undefined ? 'add' : index)).val(),
    teacher_mobile: $('#teacher_mobile_' + (index === undefined ? 'add' : index)).val()
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      $('#teacher_username_' + (index === undefined ? 'add' : index)).val(data.teacher.teacher_username);
      $('#teacher_fullname_' + (index === undefined ? 'add' : index)).val(data.teacher.teacher_fullname);
      $('#teacher_mobile_' + (index === undefined ? 'add' : index)).val(data.teacher.teacher_mobile);
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
    url: '/admin/reservation/remove',
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
    url: '/admin/reservation/cancel',
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

function setStudent(index) {
  $('body').append('\
    <div id="pop_set_student_' + index + '" class="pop_window" style="width: 50%;">\
      请输入您要制定的学生学号（必须为已注册学生）：<br>\
      <input id="student_username_' + index + '"/><br>\
      <button type="button" onclick="setStudentConfirm(' + index + ');">确认</button>\
      <button type="button" style="margin-left:20px" onclick="$(\'#pop_set_student_' + index + '\').remove();">取消</button>\
    </div>\
  ');
  optimize('#pop_set_student_' + index);
}

function setStudentConfirm(index) {
  $.post('/admin/reservation/student/set', {
    reservation_id: reservations[index].reservation_id,
    student_username: $('#student_username_' + index).val(),
  }, function(data, textStatus, xhr) {
    if (data.state == 'SUCCESS') {
      successSetStudent(index);
    } else {
      alert(data.message);
    }
  });
}

function successSetStudent(index) {
  $('#pop_set_student_' + index).remove();
  $('body').append('\
    <div id="pop_success_set_student" class="pop_window" style="width: 50%;">\
      成功指定学生！<br>\
      <button type="button" onclick="$(\'#pop_success_set_student\').remove();viewReservations();">确定</button>\
    </div>\
  ');
  optimize('#pop_success_set_student');
}

function getStudent(index) {
  $.post('/admin/reservation/student/get', {
    reservation_id: reservations[index].reservation_id
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      showStudent(data.student_info, data.reservation_info, data.student_feedback);
    } else {
      alert(data.message);
    }
  });
}

function showStudent(student, reservation, feedback) {
  $('body').append('\
    <div id="pop_show_student_' + student.id + '" class="pop_window" style="text-align: left; height: 70%; overflow:auto;">\
      <div style="width: 60%; float: left;">\
        学号：' + student.username + '<br>\
        姓名：' + student.fullname + '<br>\
        性别：' + student.gender + '<br>\
        出生年月：' + student.birthday + '<br>\
        民族：' + student.ethnic + '<br>\
        入学年份：' + student.enter_year + '<br>\
        生源地：' + student.source_place + '<br>\
        院系：' + student.college + '<br>\
        原就读学校（本科/硕士）：' + student.original_school + '<br>\
        原专业（如有转换）：' + student.original_major + '<br>\
        电子邮件：' + student.email + '<br>\
        联系电话：' + student.mobile + '<br>\
        婚姻状况：' + student.marriage + '<br>\
        健康状况：' + student.health + '<br>\
        父亲职业：' + student.father_job + '<br>\
        母亲职业：' + student.mother_job + '<br>\
        是否有兄弟姐妹：' + student.has_brother_or_sister + ' 年龄：' + student.brother_age + ' 职业：' + student.brother_job + '<br>\
        以前是否接受过职业咨询：' + student.has_career_consulting + '<br>\
        以前是否接受过心理咨询：' + student.has_mental_consulting + '<br>\
        目前是否在接受其他咨询：' + student.other_consulting_now + '<br>\
        是否有工作经验：<span id="working_experience_' + student.id + '"></span><br>\
        我们可以通过很多渠道了解与职业生涯有关的信息，最近一个月，你曾使用以下哪些方法：<br>\
        <span id="knowing_methods_' + student.id + '"></span><br><br>\
        紧急联系人：' + student.emergency_person + ' 电话：' + student.emergency_mobile + '<br>\
        <div id="student_expectation_' + student.id + '"></div>\
        <div id="student_feedback_' + student.id + '"></div>\
        <div style="margin: 10px 0">\
          <button id="export_' + student.id + '" type="button">导出</button>\
          <button type="button" onclick="$(\'#pop_show_student_' + student.id + '\').remove();">关闭</button>\
        </div>\
      </div>\
      <div style="width: 35%; float: right; border: 2px solid red; padding: 2px;">\
        <p style="margin-top: 0; background-color: red">账户相关，谨慎操作</p>\
        新密码：<input id="password_' + student.id + '" type="password"/><br>\
        确认密码：<input id="password_check_' + student.id + '" type="password"/><br>\
        <button type="button" onclick="resetStudentPassword(\'' + student.id + '\');">重置密码</button>\
        <p>删除账户</p>\
        <button type="button" onclick="deleteStudentAccount(\'' + student.id + '\');" class="btn btn-danger">删除账户</button>\
      </div>\
    </div>\
  ');
  $(function() {
    if (student.working_experience === '1') {
      $('#working_experience_' + student.id).text('有全职工作经验，工作年限：' + student.working_period);
    } else if (student.working_experience === '2') {
      $('#working_experience_' + student.id).text('有兼职工作经验或做过义工，累积工作时间：' + student.working_period);
    } else if (student.working_experience === '3') {
      $('#working_experience_' + student.id).text('没有任何工作经验');
    }
    for (var i = 0; i < student.knowing_methods.length; i++) {
      $('#knowing_methods_' + student.id).text($('#knowing_methods_' + student.id).text() + " " + knowingMethods[student.knowing_methods[i] - 1]);
    }
    if (reservation) {
      $('#student_expectation_' + student.id).append('<br>\
        此次来最主要想解决的问题是：<br><u>' + reservation.problem + '</u><br>\
        你期望职业生涯咨询帮助达到什么样的效果？<br><u>' + reservation.expectation + '</u><br>\
        期望的咨询次数约为：' + reservation.expected_time + '<br>\
        填写日期：' + reservation.time + '<br>\
      ');
    }
    $('#export_' + student.id).click(function() {
      exportStudent(student.id);
    });
    if (feedback) {

    }
  });
  optimize('#pop_show_student_' + student.id);
}

function resetStudentPassword(studentId) {
  var password = $('#password_' + studentId).val();
  var passwordConfirm = $('#password_check_' + studentId).val();
  if (password !== passwordConfirm) {
    alert('两次密码不一致，请重新输入');
    $('#password_' + studentId).val('');
    $('#password_check_' + studentId).val('');
    return;
  }
  $('body').append('\
    <div id="pop_reset" class="pop_window" style="width: 50%;">\
      您确定要重置该生的密码？<br>\
      <button type="button" onclick="$(\'#pop_reset\').remove();resetStudentPasswordConfirm(\'' + studentId + '\');">确认</button>\
      <button type="button" style="margin-left:20px" onclick="$(\'#pop_reset\').remove();">取消</button>\
    </div>\
  ');
  optimize('#pop_reset');
}

function resetStudentPasswordConfirm(studentId) {
  var password = $('#password_' + studentId).val();
  var passwordConfirm = $('#password_check_' + studentId).val();
  if (password !== passwordConfirm) {
    alert('两次密码不一致，请重新输入');
    $('#password_' + studentId).val('');
    $('#password_check_' + studentId).val('');
    return;
  }
  $.post('/admin/student/password/reset', {
    student_id: studentId,
    password: password,
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      resetStudentPasswordSuccess(studentId);
    } else {
      alert(data.message);
    }
  });
}

function resetStudentPasswordSuccess(studentId) {
  $('#password_' + studentId).val('');
  $('#password_check_' + studentId).val('');
  $('body').append('\
    <div id="pop_reset_success" class="pop_window" style="width: 50%;">\
      密码重置成功！<br>\
      <button type="button" onclick="$(\'#pop_reset_success\').remove();">确定</button>\
    </div>\
  ');
  optimize('#pop_reset_success');
}

function deleteStudentAccount(studentId) {
  $('body').append('\
    <div id="pop_delete_student_account" class="pop_window" style="width: 50%;">\
      <span style="color: red"><b>删除学生账户不可恢复，请确认操作</b></span><br>\
      <button type="button" onclick="$(\'#pop_delete_student_account\').remove();deleteStudentAccountConfirm(\'' + studentId + '\');">确认</button>\
      <button type="button" style="margin-left:20px" onclick="$(\'#pop_delete_student_account\').remove();">取消</button>\
    </div>\
  ');
  optimize('#pop_delete_student_account');
}

function deleteStudentAccountConfirm(studentId) {
  $('body').append('\
    <div id="pop_delete_student_account_confirm" class="pop_window" style="width: 50%;">\
      <span style="color: red;"><b>请再次确认删除学生账户</b></span><br>\
      <button type="button" onclick="$(\'#pop_delete_student_account_confirm\').remove();deleteStudentAccountConfirmCheck(\'' + studentId + '\');">确认</button>\
      <button type="button" style="margin-left:20px" onclick="$(\'#pop_delete_student_account_confirm\').remove();">取消</button>\
    </div>\
  ');
  optimize('#pop_delete_student_account_confirm');
}

function deleteStudentAccountConfirmCheck(studentId) {
  $.post('/admin/student/account/delete', {
    student_id: studentId,
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      deleteStudentAccountSuccess(studentId);
    } else {
      alert(data.message);
    }
  });
}

function deleteStudentAccountSuccess(studentId) {
  $('#pop_show_student_' + studentId).remove();
  $('body').append('\
    <div id="pop_delete_student_account_success" class="pop_window" style="width: 50%;">\
      学生账户已删除！<br>\
      <button type="button" onclick="$(\'#pop_delete_student_account_success\').remove();">确定</button>\
    </div>\
  ');
  optimize('#pop_delete_student_account_success');
}

function exportStudent(studentId) {
  $.post('/admin/student/export', {
    student_id: studentId,
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      window.open(data.url);
    } else {
      alert(data.message);
    }
  });
}

function exportPeriodReservations() {
  $.post('/admin/reservation/export/period', {
    from_date: $('#export_from').val(),
    to_date: $('#export_to').val(),
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      window.open(data.url);
    } else {
      alert(data.message);
    }
  });
}

Object.size = function(obj) {
  var size = 0, key;
  for (key in obj) {
    if (obj.hasOwnProperty(key)) size++;
  }
  return size;
}