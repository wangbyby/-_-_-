module GPR( input[31:0] inReg,input [4:0] rs,input [4:0] rt,input [4:0] rd,
    input flag_Write,
    input clk,
    output reg [31:0] rs1,
    output reg [31:0] rt1    
);
reg [31:0] gpr [31:0];


always @(posedge clk ) begin
    if (flag_Write) begin
      gpr[0] <= 32'b0;
        gpr[rd] <= inReg;
        rs1 <= gpr[rs]; // 可能没用
        rt1 <= gpr[rd]; // 可能没用
    end
    else begin
      gpr[0] <= 32'b0;
        rs1 <= gpr[rs];
        rt1 <= gpr[rd];
    end
end


endmodule // GPR input[31:0] 