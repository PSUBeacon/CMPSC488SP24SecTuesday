import { NavbarToggler } from '..';
import { testForChildrenInComponent, testForCustomClass, testForCustomTag, testForDefaultClass } from '../testUtils';
describe('NavbarToggler', function () {
  it('should have .navbar-toggler class', function () {
    testForDefaultClass(NavbarToggler, 'navbar-toggler');
  });
  it('should render custom tag', function () {
    testForCustomTag(NavbarToggler);
  });
  it('should render children instead of navbar-toggler-icon ', function () {
    testForChildrenInComponent(NavbarToggler);
  });
  it('should pass additional classNames', function () {
    testForCustomClass(NavbarToggler);
  });
});