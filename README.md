# Fast Dolphin TG Bot v2

**Source**:
https://github.com/polkadot21/fast_dolphin_telegram_bot

# Status

The current implementation of the `tg` bot is now the main maintained version.
The [previous version](https://github.com/polkadot21/fast_dolphin_telegram_bot
) (in Python) serves for the PoC purposes and will be deprecated after this version
implements at least the very same functionalities as the source. The further development 
will be held in this repo.

## ToDos:

1. Implement missing handlers;
2. Cover with unit tests;
3. Fix sendind training report handler;
4. Implement CI/CD (github actions).


## Environment

The current version of `golang` is `go 1.21.4`.

The environment variables required are:

    TELEGRAM_TOKEN=
    ADMIN_CHAT_ID_EVG=
    ADMIN_NAME_EVG=
    ADMIN_CHAT_ID_NAT=
    ADMIN_NAME_NAT=
    BACKEND_API=
    VERSION=
    COACH_CHAT_IDS=
    STAGE=
 

## Contribution

The contributions are very welcome.
Please submit your issues, pull requests, bug reports, or 
feature requests.