var React = require("react");
var Bootstrap = require("react-bootstrap");
var FormGroup = Bootstrap.FormGroup;
var FormControl = Bootstrap.FormControl;
var Button = Bootstrap.Button;
var Form = Bootstrap.Form;
var ControlLabel = Bootstrap.ControlLabel;
var PropTypes = React.PropTypes;


function RegisterForm(props) {
    var helpblock;
    if (props.duplicate) {
        helpblock = <HelpBlock bsClass={props.helpBlock}
                               style={{color: "red"}}>
            {props.username} is already taken.
        </HelpBlock>;
    } else {
        helpblock = <HelpBlock bsClass={props.helpBlock}
                               style={{color: "green"}}>
            {props.username} is valid.
        </HelpBlock>;
    }
    return (
        <Form onSubmit={props.onSubmitUser}>
            <FormGroup controlId="formControlsText">
                <ControlLabel>Username</ControlLabel>
                <FormControl type="text" placeholder="Enter username" onChange={props.onUpdateUser}/>
                {helpBlock}
            </FormGroup>
            <FormGroup controlId="formControlsEmail">
                <ControlLabel>Email</ControlLabel>
                <FormControl type="email" placeholder="Enter email" onChange={props.onUpdateEmail}/>
            </FormGroup>
            <FormGroup controlId="formControlsPassword">
                <ControlLabel>Password</ControlLabel>
                <FormControl type="password" onChange={props.onUpdatePassword}/>
            </FormGroup>
            <Button className={props.duplicate === true ? "disabled" : "active"} type="submit">
                Submit
            </Button>
        </Form>
    )
}

RegisterForm.propTypes = {
    onUpdateUser: PropTypes.func.isRequired,
    onSubmitUser: PropTypes.func.isRequired,
    user: PropTypes.string,
    onUpdateEmail: PropTypes.func.isRequired,
    onUpdatePassword: PropTypes.func.isRequired,
    email: PropTypes.string,
    password: PropTypes.string,
    duplicate: PropTypes.bool.isRequired,
    helpBlock: PropTypes.string.isRequired
};

module.exports = RegisterForm;
