module MUX(
    input [31:0] A,
    input [31:0] B,
    input flag,
    output [31:0] C
);
assign c = (flag)? B:A;

endmodule // MUX