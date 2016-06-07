var React = require("react");
var StudentForm = require("../components/StudentForm");

var StudentFormContainer = React.createClass({
    render: function() {
        return (
            <StudentForm />
        )
    }
});

module.exports = StudentFormContainer;