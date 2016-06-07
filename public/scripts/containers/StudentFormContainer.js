var React = require("react");
var StudentForm = require("../components/StudentForm");

var StudentFormContainer = React.createClass({
    //right now we're hardcoded with a login, but normally our state wouldn't be on first visit.
    componentWillReceiveProps: function(){
        console.log("student container get anything?", this.props);
        this.setState({
            isLoggedIn: this.props.isLoggedIn,
            username: this.props.user
        })
    },
    render: function() {
        console.log("didn't get anything?", this.props);
        return (
            <StudentForm isLoggedIn={this.props.isLoggedIn} user={this.props.user}/>
        )
    }
});

module.exports = StudentFormContainer;