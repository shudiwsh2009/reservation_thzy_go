var width = $(window).width();
var height = $(window).height();
var student;
var reservations;

function viewReservations() {
  $.getJSON('/student/reservation/view', function(json, textStatus) {
      if (json.state === 'SUCCESS') {
        console.log(json);
        student = json.student_info;
        reservations = json.reservations;
        refreshDataTable(reservations);
        optimize();
      } else {
        alert(json.message);
      }
  });
}

function refreshDataTable(reservations) {
  $('#page_maintable')[0].innerHTML = '\
    <div class="table_col" id="col_time" style="background-color:white;">\
      <div class="table_head table_cell">时间</div>\
    </div>\
    <div class="table_col" id="col_teacher" style="background-color:white;">\
      <div class="table_head table_cell">咨询师</div>\
    </div>\
    <div class="table_col" id="col_status" style="background-color:white;">\
      <div class="table_head table_cell">状态</div>\
    </div>\
    <div class="clearfix"></div>\
  ';      
  for (var i = 0; i < reservations.length; ++i) {
    $('#col_time').append('<div class="table_cell" id="cell_time_' + i + '">'
      + reservations[i].start_time.split(' ')[0] + '<br>'
      + reservations[i].start_time.split(' ')[1] + '-' + reservations[i].end_time.split(' ')[1] + '</div>');
    $('#col_teacher').append('<div class="table_cell" id="cell_teacher_' + i + '">'
      + reservations[i].teacher_fullname + '</div>');
    if (reservations[i].status === 'AVAILABLE') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i
        + '"><button type="button" id="cell_status_b_' + i + '" onclick="makeReservation(' + i
        + ')" style="padding: 2px 2px">预约</button></div>');
    } else if (reservations[i].status === 'RESERVATED') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i
        + '"><button type="button" id="cell_status_b_' + i + '" disabled="true" style="padding: 2px 2px">已预约</button>'
        + '</div>');
    } else if (reservations[i].status === 'FEEDBACK') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i
        + '"><button type="button" id="cell_status_b_' + i + '" onclick="getFeedback(' + i
        + ')" style="padding: 2px 2px">反馈</button></div>');
    } else if (reservations[i].status === 'RESERVATED_OTHER' || reservations[i].status === 'FEEDBACK_OTHER') {
      $('#col_status').append('<div class="table_cell" id="cell_status_' + i
        + '"><button type="button" id="cell_status_b_' + i + '" disabled="true" style="padding: 2px 2px">已占用</button>'
        + '</div>');
    }
  }
}

function optimize(t){
  $('#col_time').width(width * 0.45);
  $('#col_teacher').width(width * 0.25);
  $('#col_status').width(width * 0.24);
  $('#col_time').css('margin-left', width * 0.02 + 'px');

  for (var i = 0; i < reservations.length; ++i) {
    var maxHeight = Math.max(
      $('#cell_time_' + i).height(),
      $('#cell_teacher_' + i).height(),
      $('#cell_status_' + i).height()
    );

    $('#cell_time_' + i).height(maxHeight);
    $('#cell_teacher_' + i).height(maxHeight);
    $('#cell_status_' + i).height(maxHeight);
    if (i % 2 == 1) {
      $('#cell_time_' + i).css('background-color', 'white');
      $('#cell_teacher_' + i).css('background-color', 'white');
      $('#cell_status_' + i).css('background-color', 'white');
    } else {
      $('#cell_time_' + i).css('background-color', '#f3f3ff');
      $('#cell_teacher_' + i).css('background-color', '#f3f3ff');
      $('#cell_status_' + i).css('background-color', '#f3f3ff');
    }
  }
  $(t).css('left', (width - $(t).width()) / 2 - 11 + 'px');
  $(t).css('top', (height - $(t).height()) / 2 - 11 + 'px');
}

function makeReservation(index) {
  $('body').append('\
    <div class="pop_window" style="70%">\
      确定预约后请准确填写个人信息，方便职业发展中心老师与你取得联系。\
      <br>\
      <button type="button" onclick="$(\'.pop_window\').remove();makeReservationData(' + index + ');">\
        立即预约</button>\
      <button type="button" onclick="$(\'.pop_window\').remove();">暂不预约</button>\
    </div>\
  ');
  optimize('.pop_window');
}

function makeReservationData(index) {
  $('body').append('\
    <div class="pop_window" id="make_reservation_data_' + index + '" style="text-align:left; width: 90%; height: 60%; overflow: auto;">\
      <div style="text-align:center;font-size:23px">学生信息登记表</div><br>\
      ＊加粗为必填项<br>\
      <div id="student_info">\
        <span id="student_info_head" data-status="close">个人信息（点击展开）</span>\
        <div class="children">\
          <b>姓　　名：</b><input id="fullname" value="' + student.fullname + '"><br>\
          <b>性　　别：</b><select id="gender"><option value="">请选择</option><option value="男">男</option><option value="女">女</option></select><br>\
          出生年月：<input id="birthday" value="' + student.birthday + '"><br>\
          年　　龄：<input id="age" value="' + student.age + '"><br>\
          民　　族：<input id="ethnic" value="' + student.ethnic + '"><br>\
          入学年份：<input id="enter_year" value="' + student.enter_year + '"><br>\
          生源地：<input id="source_place" value="' + student.source_place + '"><br>\
          <b>院　　系：</b><input id="college" value="' + student.college + '"><br>\
          原就读学校（本科/硕士）：<input id="original_school" value="' + student.original_school + '"><br>\
          原专业（如有转换）：<input id="original_major" value="' + student.original_major + '"><br>\
          <b>电子邮件：</b><input id="email" value="' + student.email + '"><br>\
          <b>联系电话：</b><input id="mobile" value="' + student.mobile + '"><br>\
          婚姻状况：<select id="marriage"><option value="">请选择</option><option value="单身">单身</option>\
            <option value="恋爱">恋爱</option><option value="结婚">结婚</option>\
            <option value="其他">其他</option></select><br>\
          健康状况：<input id="health" value="' + student.health + '"><br>\
          父亲职业：<input id="father_job" value="' + student.father_job + '"><br>\
          母亲职业：<input id="mother_job" value="' + student.mother_job + '"><br>\
          是否有兄弟姐妹：<br>\
          <input id="has_brother_or_sister_no" name="has_brother_or_sister" type="radio" value="否">否</input><input id="has_brother_or_sister_yes" name="has_brother_or_sister" type="radio" value="是">是</input>，年龄<input id="brother_age" style="width:25px;"> 职业<input id="brother_job" style="width:35px;"><br>\
          <b>以前是否接受过职业咨询：</b><br><input id="has_career_consulting_no" name="has_career_consulting" type="radio" value="否">否</input><input id="has_career_consulting_yes" name="has_career_consulting" type="radio" value="是">是</input><br>\
          以前是否接受过心理咨询：<br><input id="has_mental_consulting_no" name="has_mental_consulting" type="radio" value="否">否</input><input id="has_mental_consulting_yes" name="has_mental_consulting" type="radio" value="是">是</input><br>\
          目前是否在接受其他咨询：<br><input id="has_other_consulting_no" name="has_other_consulting" type="radio" value="否">否</input><input id="has_other_consulting_yes" name="has_other_consulting" type="radio" value="是">是</input>，为<input id="other_consulting_now" style="width:80px;"><br>\
          是否有工作经验：<br>\
          <input id="working_experience_1" name="working_experience" type="radio" value="1">有全职工作经验，</input><input id="working_period_1" style="width:25px;">年<br>\
          <input id="working_experience_2" name="working_experience" type="radio" value="2">有兼职工作经验或做过义工，累计工作时间：</input><input id="working_period_2" style="width:50px;"><br>\
          <input id="working_experience_3" name="working_experience" type="radio" value="3">没有任何工作经验</input><br><br>\
          我们可以通过很多渠道了解与职业生涯有关的信息，最近一个月，你曾使用以下哪些方法：（可以多选）<br>\
          <input id="knowing_method_1" type="checkbox">与导师、辅导员交谈<br>\
          <input id="knowing_method_2" type="checkbox">与同学交谈<br>\
          <input id="knowing_method_3" type="checkbox">与家人交谈<br>\
          <input id="knowing_method_4" type="checkbox">与相关行业在职人员交谈<br>\
          <input id="knowing_method_5" type="checkbox">参加宣讲会、招聘会<br>\
          <input id="knowing_method_6" type="checkbox">参加相关讲座和工作坊<br>\
          <input id="knowing_method_7" type="checkbox">阅读报纸、书籍中的就业信息<br>\
          <input id="knowing_method_8" type="checkbox">浏览求职网站<br>\
          <input id="knowing_method_9" type="checkbox">选修职业辅导课程<br>\
          <input id="knowing_method_10" type="checkbox">到就业指导中心寻求帮助<br>\
          <input id="knowing_method_11" type="checkbox">向院系寻求相关就业资料<br>\
          <input id="knowing_method_12" type="checkbox">参加有关的职业生涯团体<br>\
          <input id="knowing_method_13" type="checkbox">参加过其他学生团体<br>\
          <input id="knowing_method_14" type="checkbox">在用人单位实习或者兼职<br><br>\
          <b>紧急联系人信息</b><br>\
          <b>姓　　名：</b><input id="emergency_person" value="' + student.emergency_person + '"><br>\
          <b>电　　话：</b><input id="emergency_mobile" value="' + student.emergency_mobile + '"><br>\
          <br>\
        </div>\
      </div>\
      <br>\
      <b>此次来最主要想解决的问题是什么？</b><br>\
      <textarea id="problem"></textarea><br>\
      <b>你期望职业生涯咨询帮助达到什么样的效果？</b><br>\
      <textarea id="expectation"></textarea><br>\
      期望的咨询次数约为：<br>\
      <input id="expected_time_1" type="radio">1次\
      <input id="expected_time_2" type="radio">2次\
      <input id="expected_time_3" type="radio">3次<br>\
      <button type="button" onclick="makeReservationConfirm(' + index + ');">确定</button>\
      <button type="button" onclick="$(\'.pop_window\').remove();">取消</button>\
    </div>\
  ');
  $(function() {
    $('#student_info_head').click(function() {
      if ($(this).attr('data-status') === 'close') {
          $(this).siblings().show();
          $(this).attr('data-status', 'open');
          $(this).text('个人信息（点击收起）');
      } else {
          $(this).siblings().hide();
          $(this).attr('data-status', 'close');
          $(this).text('个人信息（点击展开）');
      }
      // $(this).children('div').show().end().siblings().children('div').hide();
    });
    $('#gender').val(student.gender);
    $('#marriage').val(student.marriage);
    if (student.has_brother_or_sister === '否') {
      $('#has_brother_or_sister_no').prop('checked', true);
    } else if (student.has_brother_or_sister === '是') {
      $('#has_brother_or_sister_yes').prop('checked', true);
      $('#brother_job').val(student.brother_job);
      $('#brother_age').val(student.brother_age);
    }
    if (student.has_career_consulting === '否') {
      $('#has_career_consulting_no').prop('checked', true);
    } else if (student.has_career_consulting === '是') {
      $('#has_career_consulting_yes').prop('checked', true);
    }
    if (student.has_mental_consulting === '否') {
      $('#has_mental_consulting_no').prop('checked', true);
    } else if (student.has_mental_consulting === '是') {
      $('#has_mental_consulting_yes').prop('checked', true);
    }
    if (student.other_consulting_now === '否') {
      $('#has_other_consulting_no').prop('checked', true);
    } else if (student.other_consulting_now === '是') {
      $('#has_other_consulting_yes').prop('checked', true);
      $('#other_consulting_now').val(student.other_consulting_now);
    }
    if (student.working_experience === '1') {
      $('#working_experience_1').prop('checked', true);
      $('#working_period_1').val(student.working_period);
    } else if (student.working_experience === '2') {
      $('#working_experience_2').prop('checked', true);
      $('#working_period_2').val(student.working_period);
    } else if (student.working_experience === '3') {
      $('#working_experience_3').prop('checked', true);
    }
    for (var i = 0; i < student.knowing_methods.length; i++) {
      $('#knowing_method_' + student.knowing_methods[i]).prop('checked', true);
    }
    optimize('.pop_window');
  });
}

function makeReservationConfirm(index) {
  var fullname = $('#fullname').val();
  var gender = $('#gender').val();
  var age = $('#age').val();
  var birthday = $('#birthday').val();
  var ethnic = $('#ethnic').val();
  var enterYear = $('#enter_year').val();
  var sourcePlace = $('#source_place').val();
  var college = $('#college').val();
  var originalSchool = $('#original_school').val();
  var originalMajor = $('#original_major').val();
  var email = $('#email').val();
  var mobile = $('#mobile').val();
  var marriage = $('#marriage').val();
  var health = $('#health').val();
  var fatherJob = $('#father_job').val();
  var motherJob = $('#mother_job').val();
  var hasBrotherOrSister = $('#has_brother_or_sister_no').is(':checked') ? $('#has_brother_or_sister_no').val() 
    : ($('#has_brother_or_sister_yes').is(':checked') ? $('#has_brother_or_sister_yes').val() : "");
  var brotherAge = $('#brother_age').val();
  var brotherJob = $('#brother_job').val();
  var hasCareerConsulting = $('#has_career_consulting_no').is(':checked') ? $('#has_career_consulting_no').val()
    : ($('#has_career_consulting_yes').is(':checked') ? $('#has_career_consulting_yes').val() : "");
  var hasMentalConsulting = $('#has_mental_consulting_no').is(':checked') ? $('#has_mental_consulting_no').val()
    : ($('#has_mental_consulting_yes').is(':checked') ? $('#has_mental_consulting_yes').val() : "");
  var otherConsultingNow = $('#has_other_consulting_no').is(':checked') ? $('#has_other_consulting_no').val()
    : ($('#has_other_consulting_yes').is(':checked') ? ($('#other_consulting_now').val() ? $('#other_consulting_now').val() : $('#has_other_consulting_yes').val()) : "");
  var workingExperience = "";
  var workingPeriod = "";
  if ($('#working_experience_1').is(':checked')) {
    workingExperience = $('#working_experience_1').val();
    workingPeriod = $('#working_period_1').val();
  } else if ($('#working_experience_2').is(':checked')) {
    workingExperience = $('#working_experience_2').val();
    workingPeriod = $('#working_period_2').val();
  } else if ($('#working_experience_3').is(':checked')) {
    workingExperience = $('#working_experience_3').val();
  }
  var knowingMethods = [];
  for (var i = 1; i <= 14; i++) {
    if ($('#knowing_method_' + i).is(':checked')) {
      knowingMethods.push(i);
    }
  }
  var emergencyPerson = $('#emergency_person').val();
  var emergencyMobile = $('#emergency_mobile').val();
  var problem = $('#problem').val();
  var expectation = $('#expectation').val();
  var expectedTime = "";
  if ($('#expected_time_1').is(':checked')) {
    expectedTime = "1";
  } else if ($('#expected_time_2').is(':checked')) {
    expectedTime = "2";
  } else if ($('#expected_time_3').is(':checked')) {
    expectedTime = "3";
  }
  var payload = {
    reservation_id: reservations[index].reservation_id,
    fullname: fullname,
    gender: gender,
    college: college,
    mobile: mobile,
    email: email,
    has_career_consulting: hasCareerConsulting,
    emergency_person: emergencyPerson,
    emergency_mobile: emergencyMobile,
    age: age,
    birthday: birthday,
    ethnic: ethnic,
    enter_year: enterYear,
    source_place: sourcePlace,
    original_school: originalSchool,
    original_major: originalMajor,
    marriage: marriage,
    health: health,
    father_job: fatherJob,
    mother_job: motherJob,
    has_brother_or_sister: hasBrotherOrSister,
    brother_age: brotherAge,
    brother_job: brotherJob,
    has_mental_consulting: hasMentalConsulting,
    other_consulting_now: otherConsultingNow,
    working_experience: workingExperience,
    working_period: workingPeriod,
    knowing_methods: knowingMethods,
    problem: problem,
    expectation: expectation,
    expected_time: expectedTime,
  };
  $.ajax({
    url: '/student/reservation/make',
    type: "POST",
    dataType: 'json',
    data: payload,
    traditional: true,
  })
  .done(function(data) {
    if (data.state === 'SUCCESS') {
      makeReservationSuccess(index);
    } else {
      alert(data.message);
    }
  });
}

function makeReservationSuccess(index) {
  $('#make_reservation_data_' + index).remove();
  $('#cell_status_b_' + index).attr('disabled', 'true');
  $('#cell_status_b_' + index).text('已预约');
  $('body').append('\
    <div id="pop_make_reservation_success" class="pop_window" style="width: 50%">\
      你已预约成功，<br>\
      请关注短信提醒。<br>\
      <button type="button" onclick="$(\'#pop_make_reservation_success\').remove();viewReservations();">确定</button>\
    </div>\
  ');
  optimize('#pop_make_reservation_success');
}

function getFeedback(index) {
  $.post('/student/reservation/feedback/get', {
    reservation_id: reservations[index].reservation_id,
  }, function(data, textStatus, xhr) {
    if (data.state === 'SUCCESS') {
      showFeedback(index, data.feedback);
    } else {
      alert(data.message);
    }
  });
}

function showFeedback(index, feedback) {
  $('body').append('\
    <div class="pop_window" id="feedback_table_' + index + '"\
      style="text-align:left; width: 90%; height: 60%; overflow: auto">\
      学生反馈表<br>\
      来访者姓名：' + student.fullname + '<br>\
      第几次：<input id="feedback_consulting_count_' + index + '" value="' + feedback.consulting_count + '"><br>\
      1．我和咨询师对咨询目标的看法是一致的。<br>\
      <select id="feedback_q1_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      2．咨询师所关注的，正是我希望解决的问题。<br>\
      <select id="feedback_q2_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      3．我喜欢我的咨询师。<br>\
      <select id="feedback_q3_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      4．我主动与咨询师交流我对咨询的想法和感受。<br>\
      <select id="feedback_q4_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      5．和我的咨询师在一起，感觉很安全。<br>\
      <select id="feedback_q5_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      6．会谈中，我和咨询师一起动脑子，想办法。<br>\
      <select id="feedback_q6_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      7．我的咨询师能够接纳我哪怕是消极的想法和感受。<br>\
      <select id="feedback_q7_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      8．即使我的咨询师不赞成我所做的事情，他/她也一样关心我。<br>\
      <select id="feedback_q8_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      9．我们在咨询中所做的事情有助于实现我想要的改变。<br>\
      <select id="feedback_q9_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      10．我尝试将咨询中的收获应用于生活中。<br>\
      <select id="feedback_q10_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      11．我清楚我们要如何达成目标。<br>\
      <select id="feedback_q11_' + index + '">\
        <option value="5">总是</option>\
        <option value="4">很频繁</option>\
        <option value="3">经常</option>\
        <option value="2">有时</option>\
        <option value="1">很少</option>\
      </select><br>\
      12．总体而言，我对此次咨询感到<br>\
      <select id="feedback_q12_' + index + '">\
        <option value="5">很满意</option>\
        <option value="4">满意</option>\
        <option value="3">一般</option>\
        <option value="2">不满意</option>\
        <option value="1">很不满意</option>\
      </select><br>\
      13．总体而言，我对我的咨询师感到<br>\
      <select id="feedback_q13_' + index + '">\
        <option value="5">很满意</option>\
        <option value="4">满意</option>\
        <option value="3">一般</option>\
        <option value="2">不满意</option>\
        <option value="1">很不满意</option>\
      </select><br>\
      14.如果类似的咨询服务是收费的，你是否愿意付费？<br>\
      <select id="feedback_q14_' + index + '">\
        <option value="5">很愿意</option>\
        <option value="4">愿意</option>\
        <option value="3">说不好</option>\
        <option value="2">不愿意</option>\
        <option value="1">很不愿意</option>\
      </select><br>\
      本次会谈对你有帮助的方面是：<br>\
      <textarea id="feedback_help_' + index + '"></textarea><br>\
      本次会谈没有达到你预期的地方是：<br>\
      <textarea id="feedback_drawback_' + index + '"></textarea><br>\
      <button type="button" onclick="submitFeedback(' + index + ');">提交</button>\
      <button type="button" onclick="$(\'#feedback_table_' + index + '\').remove();">取消</button>\
    </div>\
  ');
  for (var i = 1; i <= feedback.scores.length; ++i) {
    $('#feedback_q' + i + '_' + index).val(feedback.scores[i - 1]);
  }
  $('#feedback_help_' + index).val(feedback.help);
  $('#feedback_drawback_' + index).val(feedback.drawback);
  optimize('#feedback_table_' + index);
}

function submitFeedback(index) {
  var scores = [];
  for (var i = 1; i <= 14; ++i) {
    scores.push(parseInt($('#feedback_q' + i + '_' + index).val()));
  }
  var payload = {
    reservation_id: reservations[index].reservation_id,
    consulting_count: $('#feedback_consulting_count_' + index).val(),
    scores: scores,
    help: $('#feedback_help_' + index).val(),
    drawback: $('#feedback_drawback_' + index).val(),
  };
  console.log(payload);
  $.ajax({
    url: '/student/reservation/feedback/submit',
    type: 'POST',
    dataType: 'json',
    data: payload,
    traditional: true,
  })
  .done(function(data) {
    if (data.state === 'SUCCESS') {
      successFeedback(index);
    } else {
      alert(data.message);
    }
  });
}

function successFeedback(index) {
  $('#feedback_table_' + index).remove();
  $('body').append('\
    <div id="pop_success_feedback" class="pop_window" style="width: 50%">\
      您已成功提交反馈！<br>\
      <button type="button" onclick="$(\'#pop_success_feedback\').remove();">确定</button>\
    </div>\
  ');
  optimize('#pop_success_feedback');
}