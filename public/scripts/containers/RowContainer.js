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
        console.log("componentWillreceiveProps");
        //maybe need an if statement if none of the fields are empty.
        this.setState({
            student: nextProps.student,
            course: nextProps.course,
            grade: nextProps.grade
        });
    },

    componentWillUnmount: function () {
        //when delete is fired.

    },
    render: function () {
        return (
            <Row student={this.state.grade}
                 course={this.state.course}
                 grade={this.state.grade}
                 key={this.state.id}
                 loaded={this.state.studentsLoaded}/>
        )

    }

});

module.exports = RowContainer;