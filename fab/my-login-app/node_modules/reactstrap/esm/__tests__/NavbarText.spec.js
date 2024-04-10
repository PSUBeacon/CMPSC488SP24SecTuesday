import { NavbarText } from '..';
import { testForDefaultClass, testForCustomTag, testForChildrenInComponent, testForCustomClass } from '../testUtils';
describe('NavbarText', function () {
  it('should render .navbar-text markup', function () {
    testForDefaultClass(NavbarText, 'navbar-text');
  });
  it('should render custom tag', function () {
    testForCustomTag(NavbarText);
  });
  it('should render children', function () {
    testForChildrenInComponent(NavbarText);
  });
  it('should pass additional classNames', function () {
    testForCustomClass(NavbarText);
  });
});