import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { Carousel } from '@nx-expo-monorepo/ui';

const App = () => {
  const Stack = createNativeStackNavigator();
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen
          name="Shortlink Features"
          component={() => (
            <Carousel
              title="title"
              content="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec porta leo justo, id posuere urna tempor convallis. Nulla finibus, dolor sit amet facilisis pellentesque, velit nisi tempor ipsum, nec interdum libero felis a risus. Pellentesque bibendum, dolor vel varius pulvinar, tortor leo ultrices nisi, non sodales dui quam vitae nulla. Integer sed rhoncus dui. Vestibulum bibendum diam ut leo tempus, vel vulputate magna iaculis. Suspendisse tempus magna libero, sed facilisis tellus aliquet ac. Morbi at velit ornare, posuere tortor vitae, mollis erat. Donec maximus mollis luctus. Vivamus sodales sodales dui pellentesque imperdiet. Mauris a ultricies nibh. Integer sed vehicula magna."
            />
          )}
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default App;
