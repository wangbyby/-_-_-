module IM(
    input [31:0] address,
    input clk,
    output reg [31:0] outIns
);
reg [31:0] insReg [1023:0]; 

initial 
    $readmemh("code.txt",insReg);
always @(posedge clk ) begin
    outIns <= insReg[address[9:0]];
end
endmodule // IM