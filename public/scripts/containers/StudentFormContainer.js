var React = require("react");
var StudentForm = require("../components/StudentForm");
var userFunctions = require("../utils/userFunctions");

var StudentFormContainer = React.createClass({
    //right now we're hardcoded with a login, but normally our state wouldn't be on first visit.
    getInitialState: function () {
        return {
            student: "",
            course: "",
            grade: "",
            user: ""
        };
    },
    handleUpdateStudent: function (e) {
        this.setState({
            student: e.target.value
        })
    },
    handleUpdateCourse: function (e) {
        this.setState({
            course: e.target.value
        })
    },
    handleUpdateGrade: function (e) {
        this.setState({
            grade: e.target.value
        })
    },
    handleSubmitStudent: function (e) {
        e.preventDefault();
        this.axiosAddStudent();
    },
    updateHomeContainer: function (object, studentsLoaded) {
        this.props.onStudentSubmit(object, studentsLoaded);
    },
    axiosAddStudent: function () {
        var data = {
            student: this.state.student,
            course: this.state.course,
            grade: this.state.grade
        };
        userFunctions.addStudent(data, this.state.user)
            .then(function (data) {
                this.updateHomeContainer(data.student, true);
                this.setState({
                    student: "",
                    course: "",
                    grade: ""
                })
            }.bind(this));
    },
    handlePopulate: function () {
        userFunctions.populateTable()
            .then(function (data) {
                data.student.map(function(student,index){
                    this.updateHomeContainer(student,true)
                }.bind(this))
            }.bind(this))
    },
    componentWillReceiveProps: function (nextProps) {
        this.setState({
            isLoggedIn: nextProps.isLoggedIn,
            user: nextProps.user
        })
    },
    render: function () {
        return (
            <StudentForm isLoggedIn={this.props.isLoggedIn}
                         user={this.props.user}
                         student={this.state.student}
                         course={this.state.course}
                         grade={this.state.grade}
                         onUpdateStudent={this.handleUpdateStudent}
                         onUpdateCourse={this.handleUpdateCourse}
                         onUpdateGrade={this.handleUpdateGrade}
                         onSubmitStudent={this.handleSubmitStudent}
                         onPopulate={this.handlePopulate}/>
        )
    }
});

module.exports = StudentFormContainer;