# Tanzu Plugin Architecture #

Beyond a basic set of commands provided by the core Tanzu CLI binary itself, all functionality available through the Tanzu CLI is delivered via commands provided by Tanzu Plugins.

In order for the plugins to be properly integrated with the CLI, they have to satisfy the following requirements.

1. Be executable as a standalone binary in the target platform where the CLI is
   installed.  Even though plugin commands are invoked via the CLI, the plugin itself is expected to be installed as a single executable binary.
1. It should be possible to verify the integrity and origin of the plugin
1. The plugin is expected to implement a set of standard commands listed below

## Part of the CLI-plugin contract ##

### info ###

mandatory: yes
reliance on plugin-runtime / golang CLI framework: medium

(invoked by executing the plugin binary with a single argument 'info')

This is the command through which the metadata specific to the plugin is communicated to the CLI.

```json
{
  "name": "package",
  "description": "Tanzu package management",
  "version": "v0.28.0",
  "buildSHA": "7890abcd",
  "digest": "",
  "group": "Run",
  "completionType": 0
}
```

### post-install ###

   mandatory: ?yes
   reliance on plugin-runtime / golang CLI framework: low

Need to confirm our stance on:

1. failure to find command does not block plugin install (due to needing to
   support really old plugins)
1. failure in invoking command does not block plugin install

### lint ###

   mandatory: ?yes
   reliance on plugin-runtime / golang CLI framework: high

Need to confirm how this should be used

1. Is it for certification/validation before a plugin is published? If it is
   instead also meant to be used during run time, it can be argued that the
   command is now part of the the cli-plugin contract.
1. Does it only lint for violations within the plugin?
1. What about inter-plugin inconsistencies like duplicate name/aliases

## Not Part of the CLI-plugin contract ##

### generate-docs ###

   mandatory: no
   reliance on plugin-runtime / golang CLI framework: high

Leverages cobra's built-in help-to-markdown generation.
More TBA

### version ###

   mandatory: yes
   reliance on plugin-runtime / golang CLI framework: low

Prints out the version field of the plugin-provided PluginDescriptor.

Output:
`v0.28.0-dev`

### describe ###

   mandatory: yes
   reliance on plugin-runtime / golang CLI framework: low

Merely prints out the description field of the plugin-provided PluginDescriptor.
Of minimal use today.

## Autocompletion support ##

mandatory: probably yes
reliance on plugin-runtime / golang CLI framework: high

In additional the CLI relies heavily on cobra to provide command autocompletion
support for popular shells like zsh, bash and fish by:

1. providing the `tanzu completion` command to generate the completion script
   appropriate for a shell
2. providing the __complete command to to work in conjunction with the
   completion script to provider command autocompletion UX. The completion
   experience is further extended into the plugin and their commands as well,
   as long as they are also implemented as cobra Commands
