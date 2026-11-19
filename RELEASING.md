# Releasing Gedanken

Publishing a GitHub release with a tag like `v0.1.0` runs
`.github/workflows/release.yml`. The tag sets the version and the run number
sets the build number. The workflow runs `go vet` and `go test`, builds the
Android and iOS packages and attaches them to the release. It is written by
the Giode CLI from `giode.toml`: after changing `[ci]` or a store section, run
`giode ci -f` instead of editing it, then redo the two hand edits noted in
`[ci]` (pin the CLI to `@v0.1.35-alpha.1`, `go-version-file: go.mod`).

Gedanken is a phone app (`app.platforms = ["android", "ios"]`), so it goes to
Google Play and the App Store only.

| Where | What the release gets | Needs |
| --- | --- | --- |
| Google Play | `Gedanken.aab` to upload in Play Console | a Play Console account and the upload keystore |
| GitHub, Android | `Gedanken.apk` for sideloading | nothing; signed with the upload key when the keystore secrets are set |
| App Store (iOS) | `Gedanken.ipa` to upload with Transporter | an Apple Developer account, the distribution certificate and an App Store profile |

There is no Microsoft Store or Mac App Store package: Gedanken does not target
Windows or macOS. The lessons, models and layout are made for a phone.

## Repository secrets

Set these in GitHub under Settings → Secrets and variables → Actions. Each
`*_BASE64` secret is the file encoded with `base64 -i file | pbcopy` on macOS,
or `base64 -w0 file` on Linux. Without a secret, that step is skipped and the
rest of the release still builds (unsigned AAB and APK; an iOS simulator build
instead of an `.ipa`).

| Secret | What it is |
| --- | --- |
| `ANDROID_KEYSTORE_BASE64` | The upload keystore, `release.keystore` |
| `GIODE_KEYSTORE_PASS` | Its password: signs the AAB and APK |
| `IOS_CERT_BASE64`, `IOS_CERT_PASS` | The "Apple Distribution" certificate as a `.p12`, and its password |
| `IOS_PROFILE_BASE64` | The App Store provisioning profile (`.mobileprovision`) for `com.gedanken.physics` |
| `IOS_TEAM_ID` | Your Apple team ID (10 characters, in the developer account's Membership page) |

## Google Play

1. Open a Google Play Console developer account.
2. Create the upload key, once, and keep it and its password safe (a password
   manager and an offline copy). It is not committed: `release.keystore` is in
   `.gitignore`.

   ```sh
   GIODE_KEYSTORE_PASS='a long password' giode keystore
   base64 -w0 release.keystore   # → ANDROID_KEYSTORE_BASE64
   ```

   Add `ANDROID_KEYSTORE_BASE64` and `GIODE_KEYSTORE_PASS` as secrets.
   `giode.toml` already points `android.keystore` at `release.keystore`, which
   the workflow writes from the secret.
3. In Play Console, create the app: name "Gedanken", default language English
   (United States), App, Free.
4. Publish a release on GitHub. Upload `Gedanken.aab` from it to an internal
   testing track first. Accept **Play App Signing**: Google keeps the app
   signing key and the keystore above becomes your upload key. If it is ever
   lost, Play support can reset the upload key; without Play App Signing it
   could not be replaced.
5. Fill in the listing and the App content pages with the answers below, then
   promote the release to production. A new personal developer account must
   first run a closed test with at least 12 testers for 14 days before it can
   publish to production.
6. Each upload needs a higher build number. The workflow uses the run number,
   so every release gets a new one.

**Target API level.** Google Play raises the required `targetSdkVersion` every
August. Giode targets API 35 by default. If Play Console rejects the upload for
its target API level, set `target_sdk` under `[android]` in `giode.toml` to the
level it asks for and test the app on that Android version before releasing.

## App Store (iOS)

1. Join the Apple Developer Program.
2. In Certificates, Identifiers & Profiles:
   - Register the App ID `com.gedanken.physics` (the `app.id` in `giode.toml`),
     with no extra capabilities.
   - Create an "Apple Distribution" certificate and export it with its key from
     Keychain Access as a `.p12`.
   - Create an "App Store Connect" (App Store) provisioning profile for the App ID
     with that certificate.
3. Add the four iOS secrets above.
4. In App Store Connect, create the app: iOS, name "Gedanken", primary language
   English (U.S.), bundle ID `com.gedanken.physics`, SKU `gedanken`.
5. Publish a release on GitHub. It gets `Gedanken.ipa`; upload it with the
   Transporter app, then test it in TestFlight.
6. Fill in the listing with the answers below and submit for review.

The build already carries what review checks for:

- **Icon:** `assets/icon.png`, 1024 × 1024, opaque and full-bleed.
- **Privacy manifest:** `PrivacyInfo.xcprivacy` from `[privacy]`: no tracking,
  no collected data, and the file timestamp and system boot time APIs Go uses.
- **Export compliance:** `ITSAppUsesNonExemptEncryption = false`. The app uses
  no encryption at all; it has no network code.
- **Permissions:** none, so there are no usage strings.

The app runs on iPhone and iPad (Giode builds for both), so App Store Connect
asks for iPad screenshots as well.

## Store listing

Privacy policy URL, for both stores: https://xmhrnx.github.io/gedanken/#privacy

**Name:** Gedanken

**Subtitle (App Store, 30 max):** Physics from first principles

**Short description (Google Play, 80 max):**
126 physics lessons, from motion to black holes, with the math and the history.

**Full description:**

> Gedanken is a course in physics for your phone, from motion to black holes,
> written for a curious adult who wants to understand how physicists reached
> each idea.
>
> Every lesson sets up the problem before solving it, explains the reasoning
> rather than only the result, and tells the history as a story of who was
> stuck on what.
>
> • 126 lessons in 24 chapters and 8 parts: foundations, classical mechanics,
>   matter and heat, electricity and magnetism, relativity, quantum physics,
>   atoms and nuclei, particles and cosmology
> • Real typeset mathematics, with derivations that unfold one step at a time,
>   each step explained
> • 22 interactive models: launch a projectile, shrink a star until light can't
>   leave, explore hydrogen's orbitals, set k and watch a chain reaction grow
> • History asides with the year and the place
> • A short quiz at the end of each lesson to check yourself, or skip
> • Your finished lessons are remembered, so you can see how far you've come
>
> No account, no ads, no tracking, and no internet needed: everything is in the
> app, and nothing leaves your phone.
>
> Gedanken is free and open source (Apache 2.0).

**Keywords (App Store, 100 max):**
`physics,science,relativity,quantum,black holes,mechanics,math,equations,cosmology,atom,learn,course`

**Category:** Education (both stores). App Store secondary category: Reference.

**Age rating:**
- App Store: answer "None" / "No" to every item. The Atoms and Nuclei part
  covers the history and physics of the atomic bomb as science and history,
  with no images of violence; it is not "violence" in the questionnaire's
  sense. Expected rating: 4+.
- Google Play (IARC questionnaire): category Reference, News, or Educational;
  "No" to violence, sexuality, language, controlled substances, gambling,
  user interaction and sharing of location or personal information. Expected
  rating: Everyone / PEGI 3.

**App Privacy (App Store):** "Data Not Collected". The app has no network code
and sends nothing anywhere. Tracking: No.

**Data safety (Google Play):**
- Does your app collect or share any of the required user data types? **No.**
- Is all of the user data collected by your app encrypted in transit? Not
  applicable (nothing is collected).
- Do you provide a way for users to request that their data is deleted? Not
  applicable; the only stored data, the finished-lesson list, stays on the
  device and is removed with the app.

**Google Play App content, other pages:** Ads: no ads. Target audience: 13 and
over (written for adults; choosing under-13 groups would bring in the Families
policy). Government app: no. Financial features: none. Health: none. News app:
no.

**Export compliance (App Store):** answered in the build: no encryption.

**Contact:** support URL https://github.com/xmhrnx/gedanken/issues; marketing
URL https://xmhrnx.github.io/gedanken/.

## Screenshots

Use the same five screens as the website, in this order: Home (orbital cloud
and progress), Chapter (Black Holes), Lesson (Schwarzschild's Solution),
Derivation (a static, spherical metric) and Model (the shrinking star,
`escape`). Captions are on the website.

The website's images (`site/img/screen-*.png`, 400 × 840) are too small and too
tall for the stores. Retake them from the apps:

- **App Store:** a 6.9" iPhone simulator (1320 × 2868) and a 13" iPad simulator
  (2064 × 2752); App Store Connect scales them to the smaller sizes.
- **Google Play:** at least 1080 px on the short side and no more than twice as
  tall as wide (for example 1080 × 1920 from a phone emulator); 2 to 8 phone
  screenshots. Play also needs a 512 × 512 icon (scale `assets/icon.png`) and a
  1024 × 500 feature graphic (the orbital cloud from `tools/make_icon.py` on the
  night-blue backdrop with the name works).

## Known limitations

- These files were prepared without an Android SDK or Xcode, so the AAB, APK
  and IPA have not been built yet: the first release's CI run is their first
  build. Check it on a real phone of each kind before submitting.
- Android's system backup and iCloud backups can include the finished-lesson
  list, like any app data. The privacy policy says so.
