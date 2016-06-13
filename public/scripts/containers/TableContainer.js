var React = require("react");
var Table = require("../components/Table");

//methods to add:
//check if the user is the one who made the entry, authenticate if they can edit/delete it.
//populate
//add
//delete?

var TableContainer = React.createClass({
    getInitialState: function(){
        return {
            user: "",
            isLoggedIn: false,
            student: "",
            course: "",
            grade: ""
        }
    },
    componentWillReceiveProps: function(nextProps){
        this.setState({
            user: nextProps.user,
            isLoggedIn: nextProps.isLoggedIn,
            student: nextProps.student,
            course: nextProps.course,
            grade: nextProps.grade
        })
    },
    render: function () {
        return (
            <Table user={this.state.user} student={this.state.student} course={this.state.course} grade={this.state.grade}/>
        )
    }
});

module.exports = TableContainer;