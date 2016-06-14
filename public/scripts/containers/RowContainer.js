var React = require("react");
var Row = require("../components/Row");
var userFunctions = require("../utils/userFunctions");

var RowContainer = React.createClass({
    getInitialState: function () {
        return {
            isLoggedIn: false,
            studentsLoaded: false,
            student: "",
            course: "",
            grade: "",
            id: "",
            user: ""
        }
    },
    componentWillReceiveProps: function (nextProps) {
        //maybe need an if statement if none of the fields are empty.
        this.setState({
            user: nextProps.user,
            student: nextProps.studentInfo.name,
            course: nextProps.studentInfo.course,
            grade: nextProps.studentInfo.grade,
            id: nextProps.studentInfo.id,
            studentsLoaded: nextProps.studentsLoaded
        });
    },
    render: function () {
        return (
            <Row user={this.state.user}
                student={this.state.grade}
                 course={this.state.course}
                 grade={this.state.grade}
                 key={this.state.id}
                 studentsLoaded={this.state.studentsLoaded}/>
        )

    }

});

module.exports = RowContainer;