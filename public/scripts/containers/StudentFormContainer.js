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
        //set to blank after submit.
        this.setState({
            student: "",
            course: "",
            grade: ""
        });
        this.props.onStudentSumbit(this.state.student, this.state.course, this.state.grade);

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