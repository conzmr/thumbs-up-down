import { LanguageType } from 'store/reducers/locale/langugeType'

export const environment = {
  firebase: {
    apiKey: 'AIzaSyAHOZ7rWGDODCwJMB3WIt63CAIa90qI-jg',
    authDomain: 'test-4515a.firebaseapp.com',
    databaseURL: 'https://test-4515a.firebaseio.com',
    projectId: 'test-4515a',
    storageBucket: 'test-4515a.appspot.com',
    messagingSenderId: '964743099489'
  },
  settings: {
    enabledOAuthLogin: true,
    appName: 'Thumbs Up or Down',
    appIcon: 'images/thums-up.svg',
    defaultProfileCover: 'https://www.google.com.mx/url?sa=i&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwjAvoe4_s_eAhUS2qwKHZHND44QjRx6BAgBEAU&url=https%3A%2F%2Fslate.com%2Ftechnology%2F2018%2F05%2Fthe-facebook-upvote-and-downvote-experiment-is-a-bust.html&psig=AOvVaw3ckaNq1y6mOURo_ZMemQNS&ust=1542150942946269',
    defaultLanguage: LanguageType.English
  },
  theme: {
    primaryColor: 'rgb(0, 188, 212)',
    secondaryColor: '#4d545d'
  }
}
