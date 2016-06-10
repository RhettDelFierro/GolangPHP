var React = require("react");
var StudentForm = require("../components/StudentForm");

var StudentFormContainer = React.createClass({
    //right now we're hardcoded with a login, but normally our state wouldn't be on first visit.
    getInitialState: function () {
        this.setState({
            student: "",
            course: "",
            grade: ""
        })
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
        this.setState({
            student: "",
            course: "",
            grade: ""
        })
    },
    componentWillReceiveProps: function () {
        console.log("student container get anything?", this.props);
        this.setState({
            isLoggedIn: this.props.isLoggedIn,
            username: this.props.user
        })
    },
    handleStudentSubmit: function () {

    },
    render: function () {
        console.log("didn't get anything?", this.props);
        return (
            <StudentForm isLoggedIn={this.props.isLoggedIn}
                         user={this.props.user}
                         student={this.state.student}
                         course={this.state.course}
                         grade={this.state.grade}
                         onUpdateStudent={this.handleUpdateStudent}
                         onUpdateCourse={this.handleUpdateCourse}
                         onUpdateGrade={this.handleUpdateGrade}
                         onSubmitStudent={this.handleStudentSubmit}/>
        )
    }
});

module.exports = StudentFormContainer;