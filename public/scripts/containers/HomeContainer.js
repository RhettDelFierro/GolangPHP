//will contain <StudentFormContainer> <TableContainer>
var React = require("react");
var Home = require("../components/Home");
var PropTypes = React.PropTypes;

var HomeContainer = React.createClass({
    propTypes: {
      isLoggedIn: PropTypes.bool.isRequired,
      user: PropTypes.string
    },
    getInitialState: function() {
        return {
            student: "",
            grade: "",
            course: ""
        };
    },
    handleStudentSubmit: function(student, course, grade, loaded){
        this.setState({
            student: student,
            course: course,
            grade: grade,
            studentsLoaded: loaded
        });
        //state will be updated and children will be re-rendered.
        //The table needs to get this re-rendering info. Then make an ajax call from it's componentDidMount.
        //then you're done.
    },
    render: function(){
        return (
            <Home isLoggedIn={this.props.isLoggedIn} user={this.props.user} onUpdateStudent={this.handleUpdateStudent} onStudentSubmit={this.handleStudentSubmit} student={this.state.student} course={this.state.course} grade={this.state.grade}/>
        )
    }
});

module.exports = HomeContainer;