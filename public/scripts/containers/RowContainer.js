var React = require("react");
var Row = require("../components/Row");
var userFunctions = require("../utils/userFunctions");

var RowContainer = React.createClass({
        cookieFinder: function (name) {
            var cookie = document.cookie;
            var initialValue = {};

            return cookie.split(';').reduce(function (prev, c) {
                var arr = c.split('=');
                return (arr[0].trim() == name ? arr[1] : prev);
            }, initialValue);
        },
        getInitialState: function () {
            return {
                isLoggedIn: false,
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
                student: nextProps.student,
                course: nextProps.course,
                grade: nextProps.grade
            })
        },
        componentWillMount: function () {
            //ajax call to get student info. The data is coming in from form > Home > Table > Here.
            var data = {
                student: this.state.student,
                course: this.state.course,
                grade: this.state.grade
            };
            userFunctions.addStudent(data, this.cookieFinder("Authorization"));
        },

        componentWillUnmount: function () {
            //when delete is fired.

        },
        render: function () {
            if (this.state.student && this.state.course && this.state.grade) {
                return (
                    <Row student={this.state.grade} course={this.state.course} grade={this.state.grade}
                         key={this.state.id}/>
                )
            }
        }

    })
    ;

module.exports = RowContainer;