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
    handleUpdate: function(){
        this.props.onUpdateStudent(this.refs.student.value, this.refs.course.value, this.refs.grade.value)
    },
    handleSubmitStudent: function (e) {
        e.preventDefault();
        //set to blank after submit.
        this.setState({
            student: "",
            course: "",
            grade: ""
        });
        this.props.onStudentSumbit(this.refs.studentForm.refs.student.value, this.refs.studentForm.refs.course.value, this.refs.studentForm.refs.grade.value);

        //calling userFunctions should be in  the RowContainer, not here.
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
                         onUpdate={this.handleUpdate}
                         onSubmitStudent={this.handleSubmitStudent}
            ref="studentForm"/>
        )
    }
});

module.exports = StudentFormContainer;