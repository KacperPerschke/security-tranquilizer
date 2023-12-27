# security-tranquilizer

## Context
Imagine that the security department at your job is, to put it mildly, averagely bright.

- They prevent logging in to Google and GitHub.
- If they find a fragment of code in an e-mail, they can keep the e-mail for a few days and demand an explanation why the e-mail containing the source code was sent.
- They pass images even if you report that the email looks suspicious.

## Necessity
And as a real programmer, you _think even afterhours_.
You have come up with a _part of the solution_, tested it on your private computer and would like to _send this code snippet to work so that you don't have to type it out again_ at work.

**How to do this** in the situation described above? 

## Solution
1. Use this program to encode your content in the form of an image.
1. Send image by e-mail.
1. Use this program to decode content from the image.

## The roadmap
1. [✓] Learning from mistakes.
1. [✓] Create a repo on GitHub.
1. [✓] Combine the effects of previous work with the established repo.
1. [✓] Apply "github.com/spf13/cobra"
1. [Is it done?] One file.
    1. [Probably] Be able to encode one file.
    1. [Is it done?] Be able to decode one file.
1. [Is it done?] Many files.
    1. [Is it done?] Be able to encode many files.
    1. [Is it done?] Be able to decode many files.
