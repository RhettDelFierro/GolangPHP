var React = require("react");
var Table = require("../components/Table");
var update = require("react-addons-update");

//methods to add:
//check if the user is the one who made the entry, authenticate if they can edit/delete it.
//populate
//add
//delete?

var TableContainer = React.createClass({
    getInitialState: function () {
        return {
            user: "",
            isLoggedIn: false,
            studentInfo: []
        }
    },
    componentWillReceiveProps: function (nextProps) {
        this.setState({
            user: nextProps.user,
            isLoggedIn: nextProps.isLoggedIn,
            studentInfo: update(this.state.studentInfo, {$set: nextProps.studentInfo}),
            studentsLoaded: nextProps.studentsLoaded
        })
    },
    render: function () {
        return (
            <Table user={this.state.user}
                   isLoggedIn={this.state.isLoggedIn}
                   studentInfo={this.state.studentInfo}
                   studentsLoaded={this.state.studentsLoaded}
                   onStudentDelete={this.props.onStudentDelete}/>
        )
    }
});

module.exports = TableContainer;