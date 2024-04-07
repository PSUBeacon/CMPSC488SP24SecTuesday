import React from 'react';
import {FaLock, FaLockOpen} from 'react-icons/fa';
import styled, {keyframes} from 'styled-components';

const unlockAnimation = keyframes`
  0% { transform: rotate(0deg); }
  25% { transform: rotate(-95deg); }
  75% { transform: rotate(95deg); }
  100% { transform: rotate(0deg); }
`;

const lockAnimation = keyframes`
  0% { transform: rotate(0deg); }
  25% { transform: rotate(95deg); }
  75% { transform: rotate(-95deg); }
  100% { transform: rotate(0deg); }
`;

const AnimatedIcon = styled.div`
  font-size: 2rem;
  animation: ${({isLocked}) => (isLocked ? lockAnimation : unlockAnimation)} 0.1s linear;
`;

const PadlockAnimation = ({isLocked, handleToggleLock}) => {
    return (
        <AnimatedIcon isLocked={isLocked} onClick={handleToggleLock}>
            {isLocked ? <FaLock style={{color: '#F70048'}}/> : <FaLockOpen style={{color: '#41E969'}}/>}
        </AnimatedIcon>
    );
};

export default PadlockAnimation;