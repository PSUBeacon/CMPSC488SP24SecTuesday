import React from 'react';
import { testForChildrenInComponent } from '../testUtils';
import { PopoverHeader } from '..';
describe('PopoverHeader', function () {
  it('should render children', function () {
    testForChildrenInComponent(PopoverHeader);
  });
});