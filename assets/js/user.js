var width = $(window).width();
var height = $(window).height();

function optimize(t) {
  $(t).css('left', (width - $(t).width()) / 2 - 11 + 'px');
  $(t).css('top', (height - $(t).height()) / 2 - 11 + 'px');
}

function studentLogin() {
  $.ajax({
    type: 'POST',
    async: false,
    url: '/user/student/login',
    data: {
      username: $('#username').val(),
      password: $('#password').val(),
    },
    dataType: 'json',
    success: function(data) {
      if (data.state === 'SUCCESS') {
        window.location.href = data.url;
      } else {
        alert(data.message);
      }
    }
  });
}

function studentRegister() {
  if (!$('#agree').is(':checked')) {
    alert('请先阅读并同意知情同意书');
    return;
  }
  var username = $('#username').val();
  var password = $('#password').val();
  var passwordConfirm = $('#password_confirm').val();
  if (password !== passwordConfirm) {
    alert('两次密码不一致，请重新输入');
    $('#password').val('');
    $('#password_confirm').val('');
    return;
  }
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
  if ($('#knowing_method_1').is(':checked')) {
    knowingMethods.push(1);
  } 
  if ($('#knowing_method_2').is(':checked')) {
    knowingMethods.push(2);
  } 
  if ($('#knowing_method_3').is(':checked')) {
    knowingMethods.push(3);
  } 
  if ($('#knowing_method_4').is(':checked')) {
    knowingMethods.push(4);
  } 
  if ($('#knowing_method_5').is(':checked')) {
    knowingMethods.push(5);
  } 
  if ($('#knowing_method_6').is(':checked')) {
    knowingMethods.push(6);
  } 
  if ($('#knowing_method_7').is(':checked')) {
    knowingMethods.push(7);
  } 
  if ($('#knowing_method_8').is(':checked')) {
    knowingMethods.push(8);
  } 
  if ($('#knowing_method_9').is(':checked')) {
    knowingMethods.push(9);
  } 
  if ($('#knowing_method_10').is(':checked')) {
    knowingMethods.push(10);
  } 
  if ($('#knowing_method_11').is(':checked')) {
    knowingMethods.push(11);
  } 
  if ($('#knowing_method_12').is(':checked')) {
    knowingMethods.push(12);
  } 
  if ($('#knowing_method_13').is(':checked')) {
    knowingMethods.push(13);
  }
  if ($('#knowing_method_14').is(':checked')) {
    knowingMethods.push(14);
  }
  var emergencyPerson = $('#emergency_person').val();
  var emergencyMobile = $('#emergency_mobile').val();
  var payload = {
    username: username,
    password: password,
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
  };
  console.log(payload);
  $.ajax({
    type: 'POST',
    async: false,
    url: '/user/student/register',
    data: payload,
    dataType: 'json',
    traditional: true,
    success: function(data) {
      if (data.state === 'SUCCESS') {
        window.location.href = data.url;
      } else {
        alert(data.message);
      }
    }
  });
}

function teacherLogin() {
  $.ajax({
    type: 'POST',
    async: false,
    url: '/user/teacher/login',
    data: {
      username: $('#username').val(),
      password: $('#password').val(),
    },
    dataType: 'json',
    success: function(data) {
      if (data.state === 'SUCCESS') {
        window.location.href = data.url;
      } else {
        alert(data.message);
      }
    }
  });
}

function adminLogin() {
  $.ajax({
    type: 'POST',
    async: false,
    url: '/user/admin/login',
    data: {
      username: $('#username').val(),
      password: $('#password').val(),
    },
    dataType: 'json',
    success: function(data) {
      if (data.state === 'SUCCESS') {
        window.location.href = data.url;
      } else {
        alert(data.message);
      }
    }
  });
}

function logout() {
  $.ajax({
    type: 'GET',
    async: false,
    url: '/user/logout',
    data: {},
    dataType: 'json',
    success: function(data) {
      if (data.state === 'SUCCESS') {
        window.location.href = data.url;
      }
    },
  });
}
