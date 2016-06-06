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
var RegisterFormContainer = require("../containers/RegisterFormContainer");

var styles ={
    signedIn: {
        "marginRight": 0
    }
};

function LoggedIn(props) {
    console.log("error");
    //probably just use a dropdown instead of just text.
    return (
        <Navbar.Collapse bsStyles="success">
            <Nav pullRight>
                <Navbar.Text style={styles.signedIn}>
                    Signed in as:
                </Navbar.Text>

                <NavDropdown eventKey={3} title={props.user} id="basic-nav-dropdown">
                    <MenuItem divider/>
                    <MenuItem onClick={props.onUpdateLogin}>Logout</MenuItem>
                </NavDropdown>
            </Nav>
        </Navbar.Collapse>
    )
}

function NotLoggedIn(props) {
    //remember, two different FormContainers.
    //throw in the RegisterFormContainer here?
    return (
        <Navbar.Collapse>
            <Nav>
                <Navbar.Form pullRight>
                    <FormGroup>
                        <FormControl type="text" placeholder="Search"/>
                    </FormGroup>
                    {' '}
                    <Button type="submit">Submit</Button>
                </Navbar.Form>
            </Nav>
            <Nav pullRight>
                <RegisterFormContainer />
            </Nav>
        </Navbar.Collapse>
    )
}

function LoginForm(props) {
        return props.isLoggedIn === true
            ? <LoggedIn user={props.user} onUpdateLogin={props.onUpdateLogin}/>
            : <NotLoggedIn />
}

module.exports = LoginForm;