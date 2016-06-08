//do not forget you will eventually need authentication rights for all CRUD operations other than GET.
//should the table object handle the populate() call from formObject?

var formObject = {
    student_name: "",
    student_course: "",
    student_grade: 0,
    add: function () {
        this.student_name = $('#studentName').val();
        this.student_course = $('#course').val();
        this.student_grade = $('#studentGrade').val();
        this.ajaxAdd();
        $("p").addClass("hidden").removeClass("show");
    },
    ajaxAdd: function () {
        var self = this;
        //also make the DB do this.
        var student = new AddStudent();
        student.ajaxAdd(self.student_name, self.student_course, self.student_grade);
        cancelClicked();
    },
    populate: function () {
        populate();
    },
    error: function (errorObject) {
        for (err in errorObject) {
            switch (err) {
                case "name":
                    $("#regex_name").removeClass("hidden").addClass("show").text(errorObject.name);
                case "course":
                    $("#regex_course").removeClass("hidden").addClass("show").text(errorObject.course);
                case "grade":
                    $("#regex_grade").removeClass("hidden").addClass("show").text(errorObject.grade);
                case "duplicate":
                    $("#extra_error").removeClass("hidden").addClass("show").text(errorObject.duplicate);
                //add anyway feature?
                case "database":
                    $("#extra_error").removeClass("hidden").addClass("show").text(errorObject.database);
                case "delete":
                    $("#extra_error").removeClass("hidden").addClass("show").text(errorObject.database);
            }
        }
    }
};


function Dom(name, course, grade, id) {
    var self = this;
    self.trow = $('<tr>');
    self.name = $('<td>').text(name);
    self.course = $('<td>').text(course);
    self.grade = $('<td>').text(grade);
    self.button = $('<button>').addClass("btn btn-danger").attr("id", id).on('click', function () {
        student_collection.deleteSelf($(this).attr("id"));
        // $(this).parent().remove();
        //self.arrayFunc("delete");
    }).text('Delete');
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
    deleteElement: function (id) {
        $("#" + id).parent().remove();
        //maybe will call student to delete itself too?
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
        $('.avgGrade').text(average);
    },
    handleArray: function (added) {
        this.array.push(added);
    },
    deleteSelf: function (id) {
        var self = this;
        for (i = 0; i < this.array.length; i++) {
            if (id === this.array[i].id()) {
                self.array[i].ajaxDelete();
                table.deleteElement(id);
                this.array.splice(i, 1);
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
        self.student_course = course
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
    self.ajaxAdd = function (name, course, grade) {

        var addStudent = $.ajax({
            dataType: 'json',
            data: {
                name: name,
                course: course,
                grade: grade
            },
            method: 'POST',
            url: '/api/add'
        });

        addStudent.then(function (result) {
            if (result.data) {
                self.setName(result.data.name);
                self.setCourse(result.data.course);
                self.setGrade(result.data.grade);
                self.setID(result.data.id);
                self.arrayFunc("add");
            }
            cancelClicked();

        }, function (error) {
            if (error.status === 400) {
                if (error.responseJSON.error[0] === "Records show you've already recorded this entry") {
                    var error = {
                        type: "duplicate",
                        errors: error.responseJSON.error[0]
                    };
                    ErrorHandling(error);
                } else {
                    var error = {
                        type: "regex",
                        errors: error.responseJSON.error
                    };
                    ErrorHandling(error);
                }
            }
            if (error.status === 500) {
                var error = {
                    type: "database",
                    errors: error.responseJSON.error
                };
                ErrorHandling(error);
            }
        })


    };
    self.arrayFunc = function (doing) {
        switch (doing) {
            case "add":
                student_collection.handleArray(self);
                self.addToDom();
                break;
            case "delete":
                student_collection.deleteSelf(self.id());
                break;
        }

    };
    self.addToDom = function () { //what if I made the DOM element an object itself?
        var element = new Dom(self.name(), self.course(), self.grade(), self.id());
        element.sendToTable();
    };

    //going to need a data object for logged in authentication.
    self.ajaxDelete = function () {
        var deleteStudent = $.ajax({
            dataType: 'json',
            method: 'DELETE',
            url: "/api/delete/" + self.student_id
        });

        deleteStudent.then(function (result) {
            console.log('success', result);
            if (result.data) {
                //console.log('everything is fine');
            } else {
                //console.log(result.error.responseJSON);
            }
        }, function (error) {
            if (error.status === "500") {
                console.log(error.responseJSON);
                var error = {
                    type: "delete",
                    id: error.responseJSON.error[0],
                    errors: error.resonseJSON.error[1]
                };
                ErrorHandling(error);
            } else {
                //remove the alert, and just make append "error in query" into the dom.
                alert("error in query");
            }
        })
    };
}

//put this on the form object.
function populate() {
    var populateStudents = $.ajax({
        dataType: 'json',
        method: 'GET',
        url: '/api/grades'
    });


    populateStudents.then(function (result) {
            console.log('success', result);
            console.log(result);
            if (result.data) {
                for (i = 0; i < result.data.length; i++) {
                    var student = new AddStudent();
                    student.setName(result.data[i].name);
                    student.setCourse(result.data[i].course);
                    student.setGrade(result.data[i].grade);
                    student.setID(result.data[i].id);
                    student.arrayFunc("add");
                }
                student_collection.calculateAverage()
            }
        }, function (error) {
            console.log(error.responseJSON);
            var error = {
                type: "database",
                errors: error.responseJSON.error[0]
            };
            ErrorHandling(error)
        }
    )
}

function ErrorHandling(object) {

    console.log(object);
    //just use one big switch statement.
    var errors = {};

    if (object.type == "regex") {
        for (var i = 0; i < object.errors.length; i++) {

            if (object.errors[i] === "Invalid name, please use only letters and numbers") {
                errors.name = object.errors[i];
                console.log("do we have the regex");
            }

            if (object.errors[i] === "Invalid course name, please use only letters and numbers") {
                errors.course = object.errors[i];
            }

            if (object.errors[i] === "Only numbers 0-100") {
                errors.grade = object.errors[i];
            }

        }
        formObject.error(errors);
    }

    if (object.type == "duplicate") {
        errors.duplicate = object.errors;
        formObject.error(errors);
    }

    if (object.type == "database") {
        errors.database = object.errors;
        formObject.error(errors)
    }

    if (object.type == "delete") {
        errors.delete = object.errors;
        errors.id = object.id;
        formObject.error(errors);
    }
}

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