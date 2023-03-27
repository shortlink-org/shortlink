// Handle incoming messages. Called when:
// - a message is received while the app has focus
// - the user clicks on an app notification created by a service worker
//   `messaging.onBackgroundMessage` handler.
import { getMessaging, getToken } from "firebase/messaging"
import { initializeApp } from 'firebase/app'
import { getAnalytics } from 'firebase/analytics'
import localforage from 'localforage'

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY,
  authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
  projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID,
  storageBucket: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.NEXT_PUBLIC_FIREBASE_APP_ID,
  measurementId: process.env.NEXT_PUBLIC_FIREBASE_MEASUREMENT_ID,
}

const firebaseCloudMessaging = {
  init: async () => {
    // Initialize the Firebase app with the credentials
    const app = initializeApp(firebaseConfig)
    const analytics = getAnalytics(app)

    try {
      const messaging = getMessaging()
      const tokenInLocalForage = await localforage.getItem('fcm_token')

      // Return the token if it is alredy in our local storage
      if (tokenInLocalForage !== null) {
        return tokenInLocalForage
      }

      // Request the push notification permission from browser
      const status = await Notification.requestPermission()
      if (status && status === 'granted') {
        // Get new token from Firebase

        const fcmToken = await getToken(messaging,{
          vapidKey: process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY,
        })

        // Set token in our local storage
        if (fcmToken) {
          localforage.setItem('fcm_token', fcmToken)
          return fcmToken
        }
      }
    } catch (error) {
      console.error(error)
      return null
    }
  },
}

export { firebaseCloudMessaging }
