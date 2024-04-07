import {useRef, useEffect} from 'react'

const TabTitle = ({title, prevailOnUnmount = false}) => {
    const defaultTitle = useRef(document.title);

    // eslint-disable-next-line
    useEffect(() => {
        // Set the document title to the new title
        document.title = title;

        // Cleanup function to reset the document title when the component unmounts
        return () => {
            if (!prevailOnUnmount) {
                document.title = defaultTitle.current;
            }
        };
    }, [title, prevailOnUnmount]); // Only re-run effect if title or prevailOnUnmount changes

    return null; // This component does not render anything
};

export default TabTitle;
