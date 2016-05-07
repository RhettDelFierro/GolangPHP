/**
 * Define all global variables here
 */
//var average = null;
//var student;
/**
 * student_array - global array to hold student objects
 * @type {Array}
 */
var student_array = {
    array: [],
    updateAverage: function() {
        calculateAverage(this.array);
    },
    calculateAverage: function() {
        var average = 0;
        var total_grades = 0;
        for (var i = 0; i < this.array.length; i++) {
            total_grades += parseInt(student_array[i].student_grade);
            average = Math.round(((total_grades) / (i + 1)));
        }
        console.log("average = ", average);
        $('.avgGrade').text(average);
    }
};

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
var find_form_inputs = $('.form-control:input');
/**
 * addClicked - Event Handler when user clicks the add button
 */
//***************************PT 2.
//function addClicked() {
//    console.log("this works");
//    student_name = document.getElementById("studentName").value;
//    student_course = document.getElementById("course").value;
//    student_grade = document.getElementById("studentGrade").value;
//    //student = new addStudent(student_name, student_course, student_grade);
//    var student = new addStudent(student_name, student_course, student_grade);
//    student_array.push(student);
//    addStudentAjax(student);
//    addStudentToDom(student);
//    console.log(student_array, student);
//    updateData();
//    clearAddStudentForm();
//    cancelClicked();
//}

function addClicked() {
    console.log("this works");

    //student = new addStudent(student_name, student_course, student_grade);
    var student = new addStudent(student_name, student_course, student_grade);
    // student_array.push(student);
    // addStudentAjax(student);
    student.ajax();
    //addStudentToDom(student);
    //student_array.push(student);
    addStudentToDom(student); //*****************GOLANG. Do not add to dom from array here, add from the response to avoid injection attack.

    console.log(student_array, student);
    updateData();
    clearAddStudentForm();
    cancelClicked();
}

var formObject = {
    student_name: "",
    student_course: "",
    student_grade: 0,
    add: function() {
        this.student_name = $('#studentName').val();
        this.student_course = $('#course').val();
        this.student_grade = $('#studentGrade').val();
        cancelClicked();
        this.ajax()
    },
    ajax: function() {
        var student = new AddStudent(this.student_name, this.student_course, this.student_grade);
        student.ajax();
        cancelClicked();
        console.log("may be a problem here:     ", student)
    }

};

//function addStudentAjax(student) {
//    $.ajax({
//            dataType: 'json',
//            data: {
//                name: student.student_name,
//                course: student.course,
//                grade: student.student_grade,
//                //******************************Need to not include id. Once DB sets up take this and the Golang code out.
//            },
//            method: 'POST',
//            url: '/api/add', //*****************Golang should be index.html or _tablerows.html? NO!
//            success: function (result) {
//                console.log('success!!', result);
//                if (result.success) {
//                    student.id(result.data.id);
//                    console.log('it worked man!');
//                    student_array.push(student);
//                    addStudentToDom(student);
//                } else {
//                    console.log(result.error);
//                }
//            }
//        }
//    )
//}

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
function AddStudent(name, course, grade) {
    var self = this;

    self.student_name = name,
        self.student_course = course,
        self.student_grade = grade,
        self.student_id,
    //self.delete = function () {
    //    student_array.splice(student_array.indexOf(this), 1);
    //};
    self.setName = function(name){
        self.student_name = name;
    };
    self.setCourse = function (course) {
        self.course = course
    };
    self.setGrade = function (grade) {
        self.student_grade = grade;
    };
    self.setID = function (id) {
        self.id = id;
    };
    self.name = function(){
        return self.student_name;
    };
    self.course = function () {
        return self.student_course;
    };
    self.grade = function () {
        return self.student_grade;
    };
    self.id = function () {
        return self.student_id;
    };
    self.ajax = function () {
        $.ajax({
                dataType: 'json',
                data: {
                    name: self.student_name,
                    course: self.student_course,
                    grade: self.student_grade
                },
                method: 'POST',
                url: '/api/add', //*****************Golang should be index.html or _tablerows.html? NO!
                success: function (result) {
                    console.log('success!!', result);
                    if (result) {
                        self.setID(result.id);
                        self.arrayFunc("add");
                        console.log('it worked man!');
                    } else {
                        console.log(result.error);
                    }
                }

            }
        )
    };
    self.arrayFunc = function (doing) {
        switch (doing) {
            case "add":
                student_array.push(self);
                self.addToDom();
                break;
            case "delete":
                self.ajaxDelete();
                for (i = 0; i < student_array.length; i++) {
                    if (self.student_id === student_array[i]["student_id"]) {
                        student_array[i].splice(i,1)
                    }
                }
                break;
        }
    };
    self.addToDom = function() {
        console.log("addToDom");
        var trow = $('<tr>');
        var name = $('<td>').text(self.student_name);
        var course = $('<td>').text(self.student_course);
        var grade = $('<td>').text(self.student_grade);
        var button = $('<button>').addClass("btn btn-danger").on('click', function () {
            //clearAddStudentForm();
            student_array.updateData();
            $(this).parent().remove();
            self.arrayFunc("delete");
            //need the ajax call also.
            //student.delete(); //this.delete maybe. -> NO!
        }).text('Delete');
        student.row = trow;
        //.on('click',clearAddStudentForm)
        $(trow).append(name).append(course).append(grade).append(button);
        $('tbody').append(trow);
    };
    self.ajaxDelete = function() {
        $.ajax({
            dataType: 'json',
            data: {
                //api_key: "midlWD1sMl",
                id: self.student_id
                //maybe take this info throw it into a loop into an array and post that.
            },
            method: 'POST',
            url: "/api/grades/" + self.student_id,
            success: function (result) {
                console.log('success', result);
                if (result.body) {
                    console.log('everything is fine');
                } else {
                    console.log(result.error);
                }

            }
        });
    }
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
        self.deleteStudent(student);
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
            method: 'POST',
            url: '/api/grades',
            success: function (result) {
                console.log('success', result);

                console.log(result);
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