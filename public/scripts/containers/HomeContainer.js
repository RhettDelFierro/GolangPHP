//will contain <StudentFormContainer> <TableContainer>
var React = require("react");
var Home = require("../components/Home");
var PropTypes = React.PropTypes;
var update = require("react-addons-update");

var HomeContainer = React.createClass({
    propTypes: {
        isLoggedIn: PropTypes.bool.isRequired,
        user: PropTypes.string
    },
    getInitialState: function () {
        return {
            student: "",
            grade: "",
            course: "",
            id: "",
            studentInfo: [],
            studentsLoaded: false
        };
    },
    handleStudentSubmit: function (studentObject, loaded) {
        this.setState({
            student: studentObject.name,
            course: studentObject.course,
            grade: studentObject.grade,
            id: studentObject.id,
            studentInfo: update(this.state.studentInfo, {$push: [studentObject]}),
            studentsLoaded: loaded
        });
        //state will be updated and children will be re-rendered.
        //The table needs to get this re-rendering info. Then make an ajax call from it's componentDidMount.
        //then you're done.
        console.log("in Home container: ", this.state);
    },
    handleStudentDelete: function (id) {
        var index = this.state.studentInfo.map(function (student, index) {
            if (student.id === id) {
                return index
            }
        });
        this.setState({
            studentInfo: update(this.state.studentInfo, {$splice: [[index,1]]})
        })
    },
    render: function () {
        return (
            <Home isLoggedIn={this.props.isLoggedIn} user={this.props.user} onUpdateStudent={this.handleUpdateStudent}
                  onStudentSubmit={this.handleStudentSubmit}
                  student={this.state.student}
                  course={this.state.course}
                  grade={this.state.grade}
                  studentInfo={this.state.studentInfo}
                  studentsLoaded={this.state.studentsLoaded}
                  onStudentDelete={this.handleStudentDelete}/>
        )
    }
});

module.exports = HomeContainer;