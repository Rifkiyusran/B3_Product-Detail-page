import * as React from 'react';
import { View, Text } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

const Stack = createNativeStackNavigator();

function HomeScreen() {
return (
    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
        <Text>Home Screen</Text>
    </View>
    );
}

const Rootstack = () => (
    <Stack.Navigator>
        <Stack.Screen name="Home" component={HomeScreen} />
    </Stack.Navigator>
)

export default function App() {
return (
    <NavigationContainer>
        <Rootstack/>
    </NavigationContainer>
    );
}