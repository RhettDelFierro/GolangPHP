var React = require("react");
var RegisterForm = require("../components/RegisterForm");

var RegisterFormContainer = React.createClass({
    getInitialState: function () {
        return { showModal: false };
    },

    close: function () {
        this.setState({ showModal: false });
    },
    open() {
        this.setState({ showModal: true });
    },
    render: function () {
        return (
            <RegisterForm />
        )
    }
});

module.exports = RegisterFormContainer;