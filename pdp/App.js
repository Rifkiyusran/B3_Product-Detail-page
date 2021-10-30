/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow strict-local
 */

import React, { useEffect } from 'react'; 
import Navigation from './src/Navigation/Navigation';

const App = () => {
  useEffect(() => {
    initTheme();
    SplashScreen.hide();
  }, []);

  return (
    <Navigation />
  );
};

export default App;
