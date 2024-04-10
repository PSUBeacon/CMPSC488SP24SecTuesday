import React from 'react';
import { testForChildrenInComponent } from '../testUtils';
import { PopoverBody } from '..';
describe('PopoverBody', function () {
  it('should render children', function () {
    testForChildrenInComponent(PopoverBody);
  });
});