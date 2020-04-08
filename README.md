# Create hand washing posters in local languages

Recently, [SIL International](https://www.sil.org/) launched an effort to [translate the phrase "wash your hands" into hundreds of languages](https://www.ethnologue.com/guides/health). Millions of people who speak lesser-known languages don't have a single resource on COVID-19 yet. The information gap is enormous – and lives are at stake.

"Wash your hands" isn't enough to combat this pandemic, but it's already more than some communities are hearing. On [this guide page](https://www.ethnologue.com/guides/health) you can learn how to say the phrase in 523 different languages and counting. You can also download hand-washing posters to share in print or on social media.

![SIL Health Guide Page](image.gif)

The hand washing posters were generated using the scripts in the repo. The scripts can handle a variety of complex scripts which cover almost all languages. We're releasing these bits of code here so you can float your own localized hand washing information on the "how to wash" template image and help spread the word!

## Dependencies

- [Pango](https://pango.gnome.org/)
- [Go](https://golang.org/)

## Generating posters

1. Clone this repo.

2. Create a file in [poster_text](poster_text) for each language that you want to render on a poster. Some examples are included here which illustrate a variety of complex scripts and right-to-left scripts.

3. Render the text with pango:

    ```
    $ ./generate_text.sh
    ```

4. Generate the posters:

    ```
    $ go build -o generate
    $ ./generate
    ```
