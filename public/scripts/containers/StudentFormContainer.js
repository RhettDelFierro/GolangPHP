var React = require("react");
var StudentForm = require("../components/StudentForm");
var userFunctions = require("../utils/userFunctions");

var StudentFormContainer = React.createClass({
    //right now we're hardcoded with a login, but normally our state wouldn't be on first visit.
    getInitialState: function () {
        return {
            student: "",
            course: "",
            grade: ""
        };
    },
    cookieFinder: function (name) {
        var cookie = document.cookie;
        var initialValue = {};

        return cookie.split(';').reduce(function (prev, c) {
            var arr = c.split('=');
            return (arr[0].trim() == name ? arr[1] : prev);
        }, initialValue);
    },
    handleUpdateStudent: function(e){
        this.setState({
            student: e.target.value
        })
    },
    handleUpdateCourse: function(e){
        this.setState({
            course: e.target.value
        })
    },
    handleUpdateGrade: function(e){
        this.setState({
            grade: e.target.value
        })
    },
    handleSubmitStudent: function (e) {
        e.preventDefault();
        this.axiosAddStudent();
    },
    updateHomeContainer: function(object, studentsLoaded){
      this.props.onStudentSubmit(object, studentsLoaded);
    },
    axiosAddStudent: function(){
        var data = {
            student: this.state.student,
            course: this.state.course,
            grade: this.state.grade
        };
        userFunctions.addStudent(data, this.cookieFinder("token"))
            .then(function(data){
                this.updateHomeContainer(data.student,true);
                this.setState({
                    student: "",
                    course: "",
                    grade: ""
                })
            }.bind(this));
    },
    componentWillReceiveProps: function () {
        this.setState({
            isLoggedIn: this.props.isLoggedIn,
            username: this.props.user
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
                         onSubmitStudent={this.handleSubmitStudent}/>
        )
    }
});

module.exports = StudentFormContainer;