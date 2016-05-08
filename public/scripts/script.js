/**
 * Define all global variables here
 */
//var average = null;
//var student;
/**
 * student_array - global array to hold student objects
 * @type {Array}
 */


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

var formObject = {
    student_name: "",
    student_course: "",
    student_grade: 0,
    add: function () {
        this.student_name = $('#studentName').val();
        this.student_course = $('#course').val();
        this.student_grade = $('#studentGrade').val();
        cancelClicked();
        this.ajax()
    },
    ajax: function () {
        var student = new AddStudent(this.student_name, this.student_course, this.student_grade);
        student.ajax();
        cancelClicked();
        console.log("may be a problem here:     ", student)
    }
};

var student_collection = {
    array: [],
    calculateAverage: function () {
        var average = 0;
        var total_grades = 0;
        for (var i = 0; i < this.array.length; i++) {
            total_grades += parseInt(this.array[i].student_grade);
            average = Math.round(((total_grades) / (i + 1)));
        }
        console.log("average = ", average);
        $('.avgGrade').text(average);
    },
    handleArray: function (added) {
        this.array.push(added);
    },
    deleteSelf: function (id) {
        for (i = 0; i < this.array.length; i++) {
            if (id === student_array[i].id()) {
                this.student_array[i].splice(i, 1)
            }

        }
    }
};

function cancelClicked() {
    $('input').val('');
}

function AddStudent(name, course, grade) {
    var self = this;

    self.student_name = name,
        self.student_course = course,
        self.student_grade = grade,
        self.student_id,
        //self.delete = function () {
        //    student_array.splice(student_array.indexOf(this), 1);
        //};
        self.setName = function (name) {
            self.student_name = name;
        };
    self.setCourse = function (course) {
        self.course = course
    };
    self.setGrade = function (grade) {
        self.student_grade = grade;
    };
    self.setID = function (id) {
        self.student_id = id;
    };
    self.name = function () {
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
                    if (result.data) {
                        console.log("with data:", result.data.Id);
                        self.setID(result.data.Id);
                        //self.student_id = result.data.Id;
                        self.arrayFunc("add");
                        console.log("here's self name", self.student_name);
                        console.log("here's self.student_id:", self.student_id);
                        console.log('it worked man!');
                    } else {
                        console.log(result.error);
                    }
                }

            }
        )
    };
    self.arrayFunc = function (doing, dom) {
        switch (doing) {
            case "add":
                student_collection.handleArray(self);
                self.addToDom();
                break;
            case "delete":
                console.log("arrayFunc check", dom);
                //var deleteDOM = this;
                $(dom).parent().remove();

                console.log("ajaxDelete() checked out, onto student_collection");
                student_collection.deleteSelf(self.id());
                break;
        }

    };
    self.addToDom = function () {
        console.log("addToDom");
        var trow = $('<tr>');
        var name = $('<td>').text(self.student_name);
        var course = $('<td>').text(self.student_course);
        var grade = $('<td>').text(self.student_grade);
        var button = $('<button>').addClass("btn btn-danger").on('click', function () {
            //clearAddStudentForm();
            var deleteDOM = this;
            console.log("here is the translated DOM", deleteDOM);
            if (self.ajaxDelete()) {
                self.arrayFunc("delete", deleteDOM)
            }
            // var deleteDOM = $(this).parent().remove(); //closure time?
            if (self.arrayFunc("delete")) {
                console.log("got true for arrayFunc");
                student_collection.calculateAverage()
            } //***************Golang. I don't want to delete the DOM unless I know the database delete succeeded.

            //need the ajax call also.
            //student.delete(); //this.delete maybe. -> NO!
        }).text('Delete');
        //.on('click',clearAddStudentForm)
        $(trow).append(name).append(course).append(grade).append(button);
        $('tbody').append(trow);
    };
    self.ajaxDelete = function () {
        $.ajax({
            dataType: 'json',
            data: {
                //api_key: "midlWD1sMl",
                //id: self.student_id
                //maybe take this info throw it into a loop into an array and post that.
            },
            method: 'DELETE',
            url: "/api/delete/" + self.student_id,
            success: function (result) {
                console.log('success', result);
                if (result.data) {
                    console.log('everything is fine');
                    return result.data;
                } else {
                    console.log(result.error);
                    return false
                }

            }
        });
    };
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
                    student_array.calculateAverage()
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
    cancelClicked();
    reset();
    $('th:nth-child(3)').on('click', function () {
        console.log('sort by grade');
        sortByGrade();
    });
});
/**
 * Created by Rhett on 11/9/2015.
 */