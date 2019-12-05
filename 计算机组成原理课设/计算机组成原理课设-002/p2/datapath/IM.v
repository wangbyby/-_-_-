module IM(
    input [31:0] address,
    output wire [31:0] outIns
);
// rom
reg [31:0] insReg [1023:0]; 

initial begin
    $readmemh("code.txt",insReg);
end
   assign outIns = insReg[address[11:2]];
endmodule // IM