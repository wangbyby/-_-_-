module PC( npc, flag_reset, clock, pc
);
//tmp文件夹下的 以后会删掉
input [31:2] npc;
input flag_reset;
input clock;
reg [31:0] var_tmp_reset  = 16'h0000_3000;
output [31:2] pc;

reg [31:2]reg_npc;
always @(posedge clock or flag_reset ) begin
    if (flag_reset == 1'b1)
        begin
        reg_npc = var_tmp_reset[31:2];
        pc = reg_npc;
        end
    else 
        begin
            reg_npc = npc;
            pc = npc;
        end
    end
endmodule // PC npc, flag_reset, clock, input [31:2] npc;
