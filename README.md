
# Usage
 <input>        :    input byte array
 -a <byte array>:    and <input> by <byte array>
 -o <byte array>:    or  <input> by <byte array>
 -x <byte array>:    xor <input> by <byte array>
 -n             :    not <input>
 -b             :    treat input as binary data

stdin:  input hex string
stdout: output hex string
argment: applied arithmetic hex string

echo <stdin> | alu <argument> 

read stdin
decode string to hex from <stdin> and <argument>
apply <argument> to <stdin>
encode hex array to string
write string to <stdout>

