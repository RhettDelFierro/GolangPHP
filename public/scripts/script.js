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
        var self = this;
        //also make the DB do this.
        var student = new AddStudent();
        student.ajaxAdd(self.student_name, self.student_course, self.student_grade);
        cancelClicked();
        //console.log("may be a problem here:     ", student)
    },
    populate: function () {
        populate();
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
        console.log("have the table data");
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
        console.log("average = ", average);
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
    },

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
        $.ajax({
                dataType: 'json',
                data: {
                    name: name,
                    course: course,
                    grade: grade
                },
                method: 'POST',
                url: '/api/add', //*****************Golang should be index.html or _tablerows.html? NO!
                success: function (result) {
                    console.log('success!!', result);
                    if (result.data) {
                        self.setName(result.data.name);
                        self.setCourse(result.data.course);
                        self.setGrade(result.data.grade);
                        self.setID(result.data.id);
                        //self.student_id = result.data.Id;
                        console.log("what is the student at this point:", self);
                        self.arrayFunc("add");
                    } else {
                        console.log(result.error);
                        return false;
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
                //console.log("student_colleciton:", student_collection.array);
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
                } else {
                    console.log(result.error);
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
                } else {
                    console.log(result.error);
                }
            }
        }
    )
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