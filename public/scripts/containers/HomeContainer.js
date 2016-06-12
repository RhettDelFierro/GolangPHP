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
    handleStudentSubmit: function(student, course, grade){
        this.setState({
            filterText: student,
            course: course,
            grade: grade
        });
    },
    render: function(){
        return (
            <Home isLoggedIn={this.props.isLoggedIn} user={this.props.user} onUpdateStudent={this.handleUpdateStudent} onStudentSubmit={this.handleStudentSubmit}/>
        )
    }
});

module.exports = HomeContainer;