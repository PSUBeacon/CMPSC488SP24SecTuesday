import React from 'react';
import { render, screen } from '@testing-library/react';
import { NavItem } from '..';
import { testForChildrenInComponent, testForCustomClass, testForCustomTag, testForDefaultClass } from '../testUtils';
describe('NavItem', function () {
  it('should render .nav-item class', function () {
    testForDefaultClass(NavItem, 'nav-item');
  });
  it('should render custom tag', function () {
    testForCustomTag(NavItem);
  });
  it('should render children', function () {
    testForChildrenInComponent(NavItem);
  });
  it('should pass additional classNames', function () {
    testForCustomClass(NavItem);
  });
  it('should render active class', function () {
    render( /*#__PURE__*/React.createElement(NavItem, {
      active: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('active');
  });
});