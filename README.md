[audacity](https://github.com/audacity/audacity) is a very powerful audio editing tool. 

## Plug-in
You can export a timeline or LRC text via a Nyquist plugin. Add the label plugin into your Audacity with the following steps:

> Attention: Nyquist does not support unicode characters. For example: Chinese characters.

- Open `Tools` -> `Nyquist Plug-in Installer`
- Choose `label.ny`, then confirm
- Open `Add/Remove Plug-ins`
- Find your plug-in from the new state, then enable it

How to use it?

- Open or create an Audacity project
- Add a label track from the menu `Edit` -> `Labels`, then add some text labels
- Select all tracks
- Export it from the menu `Tools` -> `Label export`

## Convert more formats
The `label` command be able to convert more formats from the Audacity label text file.

Let us export the label text file first:

- Open menu `File` -> `Export` -> `Export Labels`

then, execute the following command:

```shell
label your-label.txt
```

### Feature
- Convert to Markdown
- Convert to LRC (TODO)
