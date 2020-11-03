library(ggcorrplot)
library(reshape)
library(stringr)
library(ggplot2)
path <- "C:/Users/megha/go/src/GenerateMatrices"
files <-list.files(normalizePath(path), pattern="txt")

for (file in files) {
  if (grepl("Jaccard", file, fixed = TRUE) | grepl("Bray-Curtis", file, fixed = TRUE)) {
    table <- read.table(file.path(path, file), sep="\t", header=TRUE)
    cols <- colnames(table)
    matrix <- as.matrix(table)
    rownames(matrix) <- cols
    co=melt(matrix)
    head(co)
    ggplot(co, aes(X1, X2)) + # x and y axes => Var1 and Var2
      geom_tile(aes(fill = value)) + # background colours are mapped according to the value column
      scale_fill_gradient2(low = "#6D9EC1", 
                           mid = "white", 
                           high = "#E46726", 
                           midpoint = 0.5, limit= c(0,1.0)) + 
      theme(panel.grid.major.x=element_blank(), #no gridlines
            panel.grid.minor.x=element_blank(), 
            panel.grid.major.y=element_blank(), 
            panel.grid.minor.y=element_blank(),
            panel.background=element_rect(fill="white"), # background=white
            axis.text.x = element_text(angle=90, hjust = 1,vjust=1,size = 12,face = "bold"),
            plot.title = element_text(size=20,face="bold"),
            axis.text.y = element_text(size = 12,face = "bold")) + 
      ggtitle("Distance Heat Map") + 
      theme(legend.title=element_text(face="bold", size=14)) + 
      scale_x_discrete(name="") +
      scale_y_discrete(name="")
    imageFilename <- str_replace(file, "txt", "png")
    ggsave(imageFilename)
  }
}