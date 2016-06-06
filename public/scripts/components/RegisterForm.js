var React = require("react");
var Bootstrap = require("react-bootstrap");
var Navbar = Bootstrap.Navbar;
var Nav = Bootstrap.Nav;
var FormGroup = Bootstrap.FormGroup;
var FormControl = Bootstrap.FormControl;
var Button = Bootstrap.Button;
var NavItem = Bootstrap.NavItem;
var MenuItem = Bootstrap.MenuItem;
var NavDropdown = Bootstrap.NavDropdown;
var Form = Bootstrap.Form;
var Dropdown = Bootstrap.Dropdown;

var styles = {
    menu: {
        padding: "10px",
        minWidth: "240px"
    },
    input: {
        marginBottom: ".5em"
    },
    btn: {
        marginTop: ".75em",
        width: "100%",
        height: "32px",
        fontSize: "13px"
    }
};

function preventDefault(e) {
    e.preventDefault();
}

function DropdownForm(props) {
    //do not forget onSumbit for the Form
    return (

        <Form horizontal>
            <FormControl id="inputEmail" placeholder="Email" type="email" style={styles.input}/>
            <FormControl id="inputUsername" placeholder="Username" type="text" style={styles.input}
                         bsRole="toggle"/>
            <FormControl id="inputPassword" placeholder="Password" type="password" style={styles.input}
                         bsRole="toggle"/>
            <Button bsStyle="primary" type="submit" style={styles.btn}>Register</Button>
        </Form>

    )
}



function RegisterForm(props) {
    return (
        <NavDropdown eventKey={3} title="Register" id="basic-nav-dropdown">
            <MenuItem style={styles.menu}>
                <DropdownForm />
            </MenuItem>
        </NavDropdown>
    )
}

module.exports = RegisterForm;
