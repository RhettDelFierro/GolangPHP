var React = require("react");
var PropTypes = React.PropTypes;
var Bootstrap = require("react-bootstrap");
var LoginFormContainer = require("../containers/LoginFormContainer");
var Navbar = Bootstrap.Navbar;
var Nav = Bootstrap.Nav;
var FormGroup = Bootstrap.FormGroup;
var FormControl = Bootstrap.FormControl;
var Button = Bootstrap.Button;
var NavItem = Bootstrap.NavItem;
var MenuItem = Bootstrap.MenuItem;
var NavDropdown = Bootstrap.NavDropdown;
var ReactRouter = require("react-router");
var Link = ReactRouter.Link;

//<Nav pullRight bsStyle="pills" activeKey={1}>
//    <NavItem eventKey={1} href="/register">Register</NavItem>
//</Nav>

function RegisterToggle(props) {
    return (
        <Nav pullRight>
            <Link to="/register">
                <Button bsClass="btn btn-primary navbar-btn">Register</Button>
            </Link>
        </Nav>
    )
}


function NavigationBar(props) {
    //LoginFormContainer and RegisterFormContainer
    //<LoggedIn user={props.user} onUpdateLogin={props.onUpdateLogin}/>
    return (
        <Navbar staticTop fluid>
            <Navbar.Header>
                <Navbar.Brand>
                    <a href="#">Student Grade Table</a>
                </Navbar.Brand>
                <Navbar.Toggle />
            </Navbar.Header>
            <Navbar.Collapse>
                {!props.isLoggedIn && <RegisterToggle />}
                <LoginFormContainer isLoggedIn={props.isLoggedIn}
                                    user={props.user}
                                    onUpdateLogin={props.onUpdateLogin}/>
            </Navbar.Collapse>
        </Navbar>
    )
}

NavigationBar.propTypes = {
    isLoggedIn: PropTypes.bool.isRequired,
    user: PropTypes.string,
    onUpdateLogin: PropTypes.func.isRequired
};

module.exports = NavigationBar;