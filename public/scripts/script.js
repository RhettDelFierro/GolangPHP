/**
 * Define all global variables here
 */
var student_name = '';
var student_course = '';
var student_grade = null;
var average = null;
var student;
/**
 * student_array - global array to hold student objects
 * @type {Array}
 */
var student_array = [];

/**
 * inputIds - id's of the elements that are used to add students
 * @type {string[]}
 */
/*
 make variables storing the IDs here
 */
var find_student_name = $('#studentName');
var find_student_course = $('#course');
var find_student_grade = $('#studentGrade');
//var find_form_inputs = $('.form-control:input');
/**
 * addClicked - Event Handler when user clicks the add button
 */

function addClicked() {
    console.log("this works");
    student_name = document.getElementById("studentName").value;
    student_course = document.getElementById("course").value;
    student_grade = document.getElementById("studentGrade").value;
    //student = new addStudent(student_name, student_course, student_grade);
    var student = new addStudent(student_name, student_course, student_grade);
    student_array.push(student);
    addStudentAjax(student);
    addStudentToDom(student);
    console.log(student_array, student);
    updateData();
    clearAddStudentForm();
    cancelClicked();
}

function addStudentAjax(student) {
    $.ajax({
            dataType: 'json',
            data: {
                //api_key: "midlWD1sMl",
                name: student.student_name,
                course: student.course,
                grade: student.student_grade
                //maybe take this info throw it into a loop into an array and post that.
            },
            method: 'POST',
            //url: 'add.php',
            url: '/api/grades', //*****************Golang should be index.html or _tablerows.html?
            success: function (result) {
                console.log('success!!', result);
                if (result.success) {
                    student.id = result.data.ID;
                    console.log('it worked man!');
                    //var errors = [];
                    //errors.push(result.errors);
                    //for (i = 0; i < errors.length; i++) {
                    //    alert(errors[i]);
                    //}
                    //starting some oop.
                } else {
                    console.log(result.error);
                }
            }
            //error: function (server, timeout, request) {
            //    var different_errors = [];
            //    if (timeout) {
            //        different_errors.push(timeout);
            //    }
            //    if (server) {
            //        different_errors.push(server);
            //    }
            //    if (request) {
            //        different_errors.push(request);
            //    }
            //    alert(different_errors);
            //}
        }
    )
}

/**
 * cancelClicked - Event Handler when user clicks the cancel button, should clear out student form
 */
function cancelClicked() {
    $('input').val('');
}

/**
 * addStudent - creates a student objects based on input fields in the form and adds the object to global student array
 *
 * @return undefined
 */
function addStudent(name, course, grade, id) {
    var self = this;
    self.student_name = name,
        self.course = course,
        self.student_grade = grade,
        self.id = id,
        self.row = null;
    self.delete = function () {
        student_array.splice(student_array.indexOf(this), 1);
    };
}

/**
 * clearAddStudentForm - clears out the form values based on inputIds variable
 */
function clearAddStudentForm() {
    console.log('clearAddStudentForm');
    find_student_name.val('');
    find_student_course.val('');
    find_student_grade.val(null);
    console.log('all cleared');
    cancelClicked();
}

/**
 * calculateAverage - loop through the global student array and calculate average grade and return that value
 * @returns {number}
 */
function calculateAverage(student_array) {
    average = 0;
    var total_grades = 0;
    for (var i = 0; i < student_array.length; i++) {
        total_grades += parseInt(student_array[i].student_grade);
        average = Math.round(((total_grades) / (i + 1)));
    }
    console.log("average = ", average);
    $('.avgGrade').text(average);
    //to avoid NaN maybe add an if statement here.
    //once object is deleted from array this will work.
}

/**
 * updateData - centralized function to update the average and call student list update
 */
function updateData() {
    calculateAverage(student_array);
}

function updateStudentList() {
    var trow = $('<tr>');
    var name = $('<td>').text(student_name);
    var course = $('<td>').text(student_course);
    var grade = $('<td>').text(student_grade);
    trow.append(name).append(course).append(grade);
    // student_array.splice(this, 1); going to use the delete function.
    console.log(student_array);
}

/**
 * addStudentToDom - take in a student object, create html elements from the values and then append the elements
 * into the .student_list tbody
 * @param studentObj
 */
function addStudentToDom(student) {
    var trow = $('<tr>');
    var name = $('<td>').text(student.student_name);
    var course = $('<td>').text(student.course);
    var grade = $('<td>').text(student.student_grade);
    var button = $('<button>').addClass("btn btn-danger").on('click', function () {
        student.delete(); //this.delete maybe. -> NO!
        //clearAddStudentForm();
        updateData();
        $(this).parent().remove();
        deleteStudent(student);
    }).text('Delete');
    student.row = trow;
    //.on('click',clearAddStudentForm)
    $(trow).append(name).append(course).append(grade).append(button);
    $('tbody').append(trow);
}

function deleteStudent(student) {
    $.ajax({
        dataType: 'json',
        data: {
            //api_key: "midlWD1sMl",
            student_id: student.id
            //maybe take this info throw it into a loop into an array and post that.
        },
        method: 'POST',
        url: 'delete.php',
        success: function (result) {
            console.log('success', result);
            if (result.success) {
                console.log('everything is fine');
            } else {
                console.log(result.error);
            }

        }
    });
}

function populate() {
    $.ajax({
            dataType: 'json',
            //data: {
            //    api_key: "midlWD1sMl"
            //},
            method: 'GET',
            url: 'populate.php',
            success: function (result) {
                console.log('success', result);
                var global_result = result;
                if (global_result.success) {
                    for (i = 0; i < global_result.data.length; i++) {
                        var course = global_result.data[i].course;
                        var grade = global_result.data[i].grade;
                        var name = global_result.data[i].name;
                        var id = global_result.data[i].ID;
                        var student = new addStudent(name, course, grade, id);
                        addStudentToDom(student);
                        student_array.push(student);
                    }
                    console.log('student array before updateData(): ', student_array);
                    updateData();
                } else {
                    console.log(global_result.error);
                }
            }
        }
    )
}

/**
 * reset - resets the application to initial state. Global variables reset, DOM get reset to initial load state
 */
function reset() {
    student_array = [];
    student = {};
    student_name = '';
    student_course = '';
    student_grade = null;
    average = null;
}

/**
 * Listen for the document to load and reset the data to the initial state
 */
$(document).ready(function () {
    console.log("running!");
    clearAddStudentForm();
    reset();
    $('th:nth-child(3)').on('click', function () {
        console.log('sort by grade');
        sortByGrade();
    });
});
/**
 * Created by Rhett on 11/9/2015.
 */