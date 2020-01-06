const passport = require('passport')
const { Issuer, Strategy } = require('openid-client')

class OpenIdClient {

  constructor(baseUrl, authUrl, clientId, clientSecret, authService) {
    this.redirectUrl = `${baseUrl}/auth/callback`
    this.authUrl = authUrl
    this.clientId = clientId
    this.clientSecret = clientSecret
    this.authService = authService
  }

  async init() {
    await this.authService.setAuthClient()
    const issuer = await Issuer.discover(this.authUrl)
    this.client = new issuer.Client({
      client_id: this.clientId,
      client_secret: this.clientSecret,
      redirect_uris: [this.redirectUrl],
      response_types: ['code']
    })
    this.setupPassport()
  }

  setupPassport() {
    passport.serializeUser(function(user, cb) {
      cb(null, user)
    })

    passport.deserializeUser(function(obj, cb) {
      cb(null, obj)
    })

    const strategy = new Strategy({
      client: this.client,
      params: {
        scope: 'openid email profile',
      }
    }, function(tokenSet, cb) {
      return cb(null, tokenSet.claims())
    })
    this.strategyName = strategy.name
    passport.use(strategy)
  }
}

module.exports = OpenIdClient
