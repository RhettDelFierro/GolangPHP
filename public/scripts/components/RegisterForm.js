var React = require("react");
var Bootstrap = require("react-bootstrap");
var FormGroup = Bootstrap.FormGroup;
var FormControl = Bootstrap.FormControl;
var Button = Bootstrap.Button;
var Form = Bootstrap.Form;
var ControlLabel = Bootstrap.ControlLabel;


function RegisterForm(props) {
    //maybe wrap this in a div.
    return (
        <Form>
            <FormGroup controlId="formControlsText">
                <ControlLabel>Username</ControlLabel>
                <FormControl type="text" placeholder="Enter username" />
            </FormGroup>
            <FormGroup controlId="formControlsEmail">
                <ControlLabel>Email</ControlLabel>
                <FormControl type="email" placeholder="Enter email" />
            </FormGroup>
            <FormGroup controlId="formControlsPassword">
                <ControlLabel>Password</ControlLabel>
                <FormControl type="password" />
            </FormGroup>
            <Button type="submit">
                Submit
            </Button>
        </Form>
    )
}

module.exports = RegisterForm;
