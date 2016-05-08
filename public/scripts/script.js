var formObject = {
    student_name: "",
    student_course: "",
    student_grade: 0,
    add: function () {
        this.student_name = $('#studentName').val();
        this.student_course = $('#course').val();
        this.student_grade = $('#studentGrade').val();
        cancelClicked();
        this.ajaxAdd()
    },
    ajaxAdd: function () {
        var student = new AddStudent(); //also make the DB do this.
        student.ajax(this.student_name, this.student_course, this.student_grade);
        cancelClicked();
        console.log("may be a problem here:     ", student)
    }
};

//maybe really have a table object instead of giving the populate() method all to the form.
//make the student DOM have a method that tell the table object to do something to itself.

//dom element constructor.
function Dom(name, course, grade, id) {
    var self = this;
    self.trow = $('<tr>');
    self.name = $('<td>').text(name);
    self.course = $('<td>').text(course);
    self.grade = $('<td>').text(grade);
    self.button = $('<button>').addClass("btn btn-danger").attr("id", id).on('click', function () {
        self.ajaxDelete();
        $(this).parent().remove();
        self.arrayFunc("delete");
    }).text('Delete'),
        //.on('click',clearAddStudentForm)

        self.sendToTable = function () {
            table.makeElement(self);
        };
}

var table = {
    makeElement: function (domElement) { //should be for the table.
        $(domElement.trow).append(domElement.name).append(domElement.course).append(domElement.grade).append(domElement.button);
        $('tbody').append(domElement.trow);
        //this.array.push(domElement)
    },
    deleteElement: function(id) {
        $("#" + id).parent().remove();
        //maybe will calll student to delete itself too?
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
            if (id === this.array[i].id()) {
                table.deleteElement(id);
                this.array.splice(i, 1)
            }

        }
        this.calculateAverage()
    }
};

function cancelClicked() {
    $('input').val('');
}

function AddStudent() {
    var self = this;

    self.student_name = "";
    self.student_course = "";
    self.student_grade = 0;
    self.student_id = "";
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
    self.ajax = function (name, grade, course) {
        $.ajax({
                dataType: 'json',
                data: {
                    name: name,
                    course: grade,
                    grade: course
                },
                method: 'POST',
                url: '/api/add', //*****************Golang should be index.html or _tablerows.html? NO!
                success: function (result) {
                    console.log('success!!', result);
                    if (result.data) {
                        self.setName(result.data.name);
                        self.setName(result.data.course);
                        self.setName(result.data.grade);
                        self.setID(result.data.id);
                        //self.student_id = result.data.Id;
                        self.arrayFunc("add");
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
                student_collection.handleArray(self);
                self.addToDom();
                console.log("student_colleciton:", student_collection.array);
                break;
            case "delete":
                console.log("ajaxDelete() checked out, onto student_collection");
                student_collection.deleteSelf(self.id());
                break;
        }

    };
    self.addToDom = function () { //what if I made the DOM element an object itself?
        var element = new Dom(self.name(), self.course(), self.grade(), self.id());
        element.sendToTable();
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